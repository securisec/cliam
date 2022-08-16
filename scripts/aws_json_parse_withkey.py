# valid request is in operations
# request time is in shapes
import json
from pathlib import Path
import re
import pyperclip

def toSnakeCase(s):
    return re.sub(r'(?<!^)(?=[A-Z])', '_', s).lower()

def getPathParam(s, p):
    return re.sub(r'\{.+\}', f'{{{{.{toSnakeCase(p)}}}}}', s)

def postJsonEndpoint(operation: str, param: str, jsonVersion: str) -> str:
    return f""""{operation}": {{
		Method: "POST",
		Headers: map[string]string{{
			shared.CONTENT_TYPE_HEADER: {'aws_JSON_1_1' if jsonVersion == '1.1' else 'aws_JSON_1_0'},
			aws_X_AMZ_TARGET:           "{data['metadata']['targetPrefix']}.{operation}",
		}},
		Permission: "{operation}",
        IsExtra: true,
        ExtraComponentBodyKey:  "{param}",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "{toSnakeCase(param)}",
}},"""

def getRequestEndpoint(operation: str, param: str) -> str:
    return f""""{operation}": {{
        ServiceSuffix:          "{getPathParam(v['http']['requestUri'], param)}",
		Permission:             "{operation}",
		ExtraComponentLocation: "path",
		IsExtra:                true,
        ExtraCommandLineFlag:   "{toSnakeCase(param)}",
}},"""

def postFormEndpoint(operation: str, param: str) -> str:
    return f""""{operation}": {{
		Method: "POST",
		FormData: map[string]string{{
			"Action":  "{operation}",
			"Version": "{data['metadata']['apiVersion']}",
		}},
		Headers: map[string]string{{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		}},
		Permission: "{operation}",
        IsExtra: true,
        ExtraComponentBodyKey:  "{param}",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "{toSnakeCase(param)}",
}},"""



fileName = '../temp/awsapis/cognito-sync-2014-06-30.normal.json'

path = Path.cwd() / 'iam-enumerator' / fileName
with open(str(path.resolve()), 'r') as f:
    data = json.loads(f.read())

out = ['// extra']
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
            method = v['http']['method']
            param = shape['required'][0]
            if method == 'GET':
                out.append(getRequestEndpoint(operation, param))

            if method == 'POST':
                isFormEncoded = 'targetPrefix' not in data['metadata']
                if isFormEncoded:
                    out.append(postFormEndpoint(operation, param))
                else:
                    out.append(postJsonEndpoint(operation, param, data['metadata']['jsonVersion']))

    except Exception as e:
        # print(operation)
        raise e
        # print(e)

x = '\n'.join(out)
# pyperclip.copy(x)
print(x)