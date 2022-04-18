# valid request is in operations
# request time is in shapes
from fileinput import filename
import json

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



fileName = '/private/tmp/aws-sdk-js/apis/databrew-2017-07-25.normal.json'
with open(fileName, 'r') as f:
    data = json.loads(f.read())

for operation, v in data['operations'].items():
    # print(operations)
    try:
        if f'{operation}Request' in data['shapes']:
            shape = data['shapes'][f'{operation}Request']
        else:
            shape = data['shapes'][f'{operation}Input']
        if 'required' in shape:
            continue

        http = v['http']
        endpoint = http['requestUri'][1:]
        if http['method'] == 'POST':
            if http['requestUri'] == '/':
                print(postRequests(operation, data['metadata']['targetPrefix']))
            else:
                print(postRequestEndpoint(endpoint, operation))
        elif http['method'] == 'GET':
            print(getRequestEndpoint(endpoint, operation))
        
    except Exception as e:
        print(e)