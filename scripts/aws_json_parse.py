# valid request is in operations
# request time is in shapes
from fileinput import filename
import json
from pathlib import Path

def postRequests(action: str, prefix: str) -> str:
    return f"""{{
    Method: "POST",
    JsonData: `{{}}`,
    Headers: map[string]string{{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "{prefix}.{action}",
		}},
    Permission: "{action}",
}},"""

def postRequestEndpoint(suffix: str, action: str) -> str:
    return f"""{{
    Method: "POST",
    ServiceSuffix: "{suffix}",
    JsonData: `{{}}`,
    Headers: map[string]string{{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		}},
    Permission: "{action}",
}},"""

def getRequestEndpoint(suffix: str, action: str) -> str:
    return f"""{{
    Method: "GET",
    ServiceSuffix: "{suffix}",
    Permission: "{action}",
}},"""

def postFormEndpoint(version: str, action: str) -> str:
    return f"""{{
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



fileName = '../temp/awsapis/iot-2015-05-28.normal.json'
path = Path.cwd() / 'iam-enumerator' / fileName
with open(str(path.resolve()), 'r') as f:
    data = json.loads(f.read())

for operation, v in data['operations'].items():
    # print(operations)
    try:
        if 'input' in v:
            shape = data['shapes'][v['input']['shape']]
            if 'required' in shape:
                continue

        http = v['http']
        endpoint = http['requestUri'][1:]
        isFormEncoded = 'targetPrefix' not in data['metadata']
        if http['method'] == 'POST':
            if http['requestUri'] == '/':
                if isFormEncoded:
                    print(postFormEndpoint(data['metadata']['apiVersion'], operation))
                else:
                    print(postRequests(operation, data['metadata']['targetPrefix']))
            else:
                print(postRequestEndpoint(endpoint, operation))
        elif http['method'] == 'GET':
            print(getRequestEndpoint(endpoint, operation))
        
    except Exception as e:
        print(operation)
        raise e
        print(e)