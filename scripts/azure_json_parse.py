from fileinput import filename
import json
from pathlib import Path
import re
import pyperclip


re_param = re.compile(r"{(.*?)}")


def process_parameters(oper):
    params = []
    if not "parameters" in oper:
        return params
    for p in oper["parameters"]:
        if "$ref" in p:
            continue
        if p["in"] != "path":
            continue
        if p["required"]:
            params.append(p)
    return params


def convert_path(path):
    return re_param.sub("{{.\g<1>}}", path)


def process_get(path, oper, method):
    return f"""{{
    Path: "{convert_path(path)}",
	Method: "{method.upper()}",
	QueryValues:   map[string]string{{
        "api-version": "{VERSION}",
    }},
	OperationID:    "{oper['operationId']}",
    Resource:       "{RESOURCE}",
}},"""


def buildPolicy(path):
    hold = []
    with path.resolve().open("r") as f:
        data = json.loads(f.read())

        for endpoint in data["paths"].keys():
            for method, resource in data["paths"][endpoint].items():
                if method not in ["get", "post"]:
                    continue
                params = process_parameters(resource)
                if method == "get" or method == "post":
                    hold.append(process_get(endpoint, resource, method))
    return hold


def getPolicies(resource, specification, version):
    dirpath = Path(
        f"temp/azure-rest-api-specs/specification/{specification}/resource-manager/{resource}/stable/{version}/"
    )
    return [f"../{x}" for x in dirpath.glob("*.json")]


# pyperclip.copy(o)

SPECIFICATION = "apimanagement"
RESOURCE = "Microsoft.ApiManagement"
VERSION = "2021-08-01"


for resource_path in getPolicies(RESOURCE, SPECIFICATION, VERSION):

    path = Path("cliam/" + resource_path)
    resource = path.stem

    save_path = Path(f"azure/policy/{RESOURCE}.{resource}.go")
    if save_path.exists():
        continue

    policies = buildPolicy(path)
    var_name = RESOURCE.replace(".", "_") + f"_{resource}"

    template = f"""package policy

var {var_name} = []Policy{{
    {''.join(policies)}
}}
"""

    save_path.write_text(template)

    print(f'"{RESOURCE}.{resource}": policy.{var_name},')
