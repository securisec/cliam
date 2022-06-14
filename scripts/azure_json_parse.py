from fileinput import filename
import json
from pathlib import Path
import re
import pyperclip

path = Path(
    "temp/azure-rest-api-specs/specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/StaticSites.json"
)
VERSION = "2021-03-01"
RESOURCE = "Microsoft.Web"

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


def process_get(path, oper, method, extra: bool):
    return f"""{{
    Path: "{convert_path(path)}",
	Method: "{method.upper()}",
	QueryValues:   map[string]string{{
        "api-version": "{VERSION}",
    }},
	OperationID:    "{oper['operationId']}",
    Resource:       "{RESOURCE}",
    {"IsExtra: true," if extra else ""}
    {'ExtraLocation: "path",' if extra else ""}
}},"""


def process_get_extra(path, oper):
    # TODO 🚧
    pass


hold = []
with path.resolve().open("r") as f:
    data = json.loads(f.read())

    for endpoint in data["paths"].keys():
        for method, resource in data["paths"][endpoint].items():
            if method not in ["get", "post"]:
                continue
            params = process_parameters(resource)
            if method == "get" or method == "post":
                hold.append(process_get(endpoint, resource, method, len(params) != 0))

print(len(hold))
o = "\n".join(hold)
pyperclip.copy(o)
