# valid request is in operations
# request time is in shapes
from fileinput import filename
import json
from pathlib import Path
from typing import List, Any


def postRequests(action: str, prefix: str) -> str:
    return f""""{action}": {{
    Method: "POST",
    JsonData: map[string]string{{}},
    Headers: map[string]string{{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "{prefix}.{action}",
		}},
    Permission: "{action}",
}},"""


def postRequestEndpoint(suffix: str, action: str) -> str:
    return f""""{action}":  {{
    Method: "POST",
    ServiceSuffix: "{suffix}",
    JsonData: map[string]string{{}},
    Headers: map[string]string{{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		}},
    Permission: "{action}",
}},"""


def getRequestEndpoint(suffix: str, action: str) -> str:
    return f""""{action}":  {{
    Method: "GET",
    ServiceSuffix: "{suffix}",
    Permission: "{action}",
}},"""


def postFormEndpoint(version: str, action: str) -> str:
    return f""""{action}":  {{
    Method: "POST",
		FormData: map[string]string{{
			"Action":  "{action}",
			"Version": "{version}",
		}},
		Headers: map[string]string{{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		}},
		Permission: "{action}",
}},"""


def get_policies(filename: str, ignore_creates_deletes: bool = False) -> List[Any]:
    # TODO 🔥 ignore creates and deletes
    out = []
    path = Path.cwd() / "iam-enumerator" / filename.replace("../", "", 1)
    with open(str(path.resolve()), "r") as f:
        data = json.loads(f.read())

    for operation, v in data["operations"].items():
        # skip create operations
        if ignore_creates_deletes and operation.startswith('Create'):
            continue
        # print(operations)
        try:
            if "input" in v:
                shape = data["shapes"][v["input"]["shape"]]
                if "required" in shape:
                    continue

            http = v["http"]
            endpoint = http["requestUri"][1:]
            isFormEncoded = "targetPrefix" not in data["metadata"]
            if http["method"] == "POST":
                if http["requestUri"] == "/":
                    if isFormEncoded:
                        out.append(
                            postFormEndpoint(data["metadata"]["apiVersion"], operation)
                        )
                    else:
                        out.append(
                            postRequests(operation, data["metadata"]["targetPrefix"])
                        )
                else:
                    out.append(postRequestEndpoint(endpoint, operation))
            elif http["method"] == "GET":
                out.append(getRequestEndpoint(endpoint, operation))

        except Exception as e:
            print(operation)
            raise e

    return out
