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



fileName = '../temp/awsapis/ec2-2016-11-15.normal.json'
path = Path.cwd() / 'iam-enumerator' / fileName
with open(str(path.resolve()), 'r') as f:
    data = json.loads(f.read())

for operation, v in data['operations'].items():
    # print(operations)
    try:
        if 'input' in v:
            if not any([operation.startswith(x) for x in ['Get', 'List', 'Describe']]):
                continue
            # if 'shape' not in v['input'] or 'required' not in v['input']['shape']:
            #     continue
            shape = data['shapes'][v['input']['shape']]
            if not 'required' in shape:
                continue
            if not len(shape['required']) == 1:
                continue
            if not shape['required'][0] == 'InstanceId':
                continue
            print(f"""{{
		Method: "POST",
		FormData: map[string]string{{
			"Action":  "{operation}",
			"Version": "2016-11-15",
		}},
		Headers: map[string]string{{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		}},
		Permission: "{operation}",
        IsExtra: true,
}},""")

        # http = v['http']
        # endpoint = http['requestUri'][1:]
        # isFormEncoded = 'targetPrefix' not in data['metadata']
        # if http['method'] == 'POST':
        #     if http['requestUri'] == '/':
        #         if isFormEncoded:
        #             print(postFormEndpoint(data['metadata']['apiVersion'], operation))
        #         else:
        #             print(postRequests(operation, data['metadata']['targetPrefix']))
        #     else:
        #         print(postRequestEndpoint(endpoint, operation))
        # elif http['method'] == 'GET':
        #     print(getRequestEndpoint(endpoint, operation))
        
    except Exception as e:
        # print(operation)
        raise e
        # print(e)


# s3
# if 'input' in v:
#             if 'Create' in operation or 'Delete' in operation or 'Put' in operation:
#                 continue
#             shape = data['shapes'][v['input']['shape']]
#             if not len(shape['required']) == 1:
#                 continue
#             print(f"""{{
#                 ServiceSuffix:          "{v['http']['requestUri'].replace('{Bucket}', '{{.}}')}",
# 		Permission:             "{operation}",
# 		ExtraComponentLocation: "path",
# 		IsExtra:                true,
# }},""")