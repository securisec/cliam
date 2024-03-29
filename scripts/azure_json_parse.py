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
                # params = process_parameters(resource)
                if method == "get" or method == "post":
                    hold.append('"{}": {}'.format(resource['operationId'], process_get(endpoint, resource, method)))

    return '\n'.join(hold)


def getPolicies(resource, specification, version):
    dirpath = Path(
        f"temp/azure-rest-api-specs/specification/{specification}/resource-manager/{resource}/stable/{version}/"
    )
    return [f"../{x}" for x in dirpath.glob("*.json")]


# pyperclip.copy(o)

all_services = [dirs.stem for dirs in Path('temp/azure-rest-api-specs/specification/').glob('*') if dirs.is_dir()]
print(all_services)

for SPECIFICATION in all_services:
    # SPECIFICATION = "keyvault"

    try:
        rdirs =  [dirs for dirs in Path(f'temp/azure-rest-api-specs/specification/{SPECIFICATION}/resource-manager/').glob("*") if dirs.is_dir()]

        for dirs in rdirs:
            RESOURCE = dirs.name
            VERSION = [x for x in sorted((dirs / "stable").glob("*"))][-1].name


            for resource_path in getPolicies(RESOURCE, SPECIFICATION, VERSION):

                path = Path("cliam/" + resource_path)
                resource = path.stem

                save_path = Path(f"azure/policy/{RESOURCE}.{resource}.go")
                print(save_path)
                # if save_path.exists():
                #     continue

                policies = buildPolicy(path)
                var_name = RESOURCE.replace(".", "_") + f"_{resource}"

                template = f"""package policy

    // {var_name} policy
    //
    var {var_name} = map[string]Policy{{
        {policies}
    }}
    """

                save_path.write_text(template)
                # print(policies)

    except:
        continue
