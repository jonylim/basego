define({ "api": [
  {
    "type": "post",
    "url": "/v1/account/profile/accept_tos",
    "title": "Accept Terms of Service",
    "version": "1.2.2",
    "name": "AcceptTOS",
    "group": "AccountAPI",
    "permission": [
      {
        "name": "account"
      }
    ],
    "description": "<p>Accept Terms of Service.</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "long",
            "optional": false,
            "field": "createdTime",
            "description": "<p>The account's created time, to validate the request.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Request Example:",
          "content": "{\n  \"createdTime\": 1566452967522\n}",
          "type": "json"
        }
      ]
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "boolean",
            "optional": false,
            "field": "success",
            "description": "<p>If accepted successfully.</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "message",
            "description": "<p>The message.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success Response:",
          "content": "HTTP/1.1 200 OK\n{\n  \"status\": 200,\n  \"error\": {\n    \"code\": \"\",\n    \"message\": \"\",\n    \"field\": \"\"\n  },\n  \"data\": {\n    \"success\": true,\n    \"message\": \"Terms of Service accepted\",\n  }\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "ParamValidationFailed",
            "description": "<p>The parameter validation failed.</p>"
          },
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "HeaderValidationFailed",
            "description": "<p>The request header validation failed.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "ParamValidationFailed:",
          "content": "HTTP/1.1 200 OK\n{\n  \"status\": 400,\n  \"error\": {\n    \"code\": \"40002\",\n    \"message\": \"Created time is invalid\",\n    \"field\": \"createdTime\"\n  },\n  \"data\": {}\n}",
          "type": "json"
        },
        {
          "title": "HeaderValidationFailed:",
          "content": "HTTP/1.1 200 OK\n{\n  \"status\": 400,\n  \"error\": {\n    \"code\": \"40001\",\n    \"message\": \"Request headers are required (API-Key, Authorization, Device-Platform)\",\n    \"field\": \"\"\n  },\n  \"data\": {}\n}",
          "type": "json"
        }
      ]
    },
    "filename": "internal/app/basego-api/v1/endpoint/accountapi/profile-accept-tos.go",
    "groupTitle": "Account API",
    "groupDescription": "<h4>HTTP Request Headers</h4> <table> <thead> <tr> <th><strong>Header Name</strong></th> <th style=\"text-align:center\"><strong>Required</strong></th> <th><strong>Description</strong></th> </tr> </thead> <tbody> <tr> <td>API-Key</td> <td style=\"text-align:center\">✓</td> <td>API key for accessing the API.</td> </tr> <tr> <td>App-Identifier</td> <td style=\"text-align:center\"></td> <td>The app's identifier (package name for Android, bundle ID for iOS, or origin URL for web).</td> </tr> <tr> <td>Authorization</td> <td style=\"text-align:center\">✓</td> <td>Access token to validate the user session.<br>Format: <code>Bearer <i>&lt;access_token&gt;</i></code></td> </tr> <tr> <td>Content-Type</td> <td style=\"text-align:center\"></td> <td>Content type of the request body.</td> </tr> <tr> <td>Device-Identifier</td> <td style=\"text-align:center\">✓</td> <td>The device ID (optional for web).</td> </tr> <tr> <td>Device-Model</td> <td style=\"text-align:center\">✓</td> <td>Model name of the device (optional for web).</td> </tr> <tr> <td>Device-Platform</td> <td style=\"text-align:center\">✓</td> <td>The device's platform. Values are <code>android</code>, <code>ios</code>, or <code>web</code>.</td> </tr> <tr> <td>User-Agent</td> <td style=\"text-align:center\"></td> <td>The user agent of the client accessing the API.</td> </tr> </tbody> </table> <h4>HTTP Response Status Codes</h4> <table> <thead> <tr> <th style=\"text-align:center\"><strong>Code</strong></th> <th><strong>Description</strong></th> </tr> </thead> <tbody> <tr> <td style=\"text-align:center\">200</td> <td>OK, request proceed without error.</td> </tr> <tr> <td style=\"text-align:center\">204</td> <td>Request proceed successfully and is not returning any content (empty <code>data</code>).</td> </tr> <tr> <td style=\"text-align:center\">400</td> <td>Bad request, either the request header or the API parameter validation failed.</td> </tr> <tr> <td style=\"text-align:center\">401</td> <td>Unauthorized access, either the HTTP request header <code>Authorization</code> is missing or invalid.</td> </tr> <tr> <td style=\"text-align:center\">403</td> <td>Forbidden, the user does not have access to the requested resource.</td> </tr> <tr> <td style=\"text-align:center\">404</td> <td>Not found, requested resource does not exist.</td> </tr> <tr> <td style=\"text-align:center\">429</td> <td>Too many requests sent in a given amount of time, intended for use with rate-limiting schemes.</td> </tr> <tr> <td style=\"text-align:center\">491</td> <td>The API key specified in HTTP request header <code>API-Key</code> is invalid.</td> </tr> <tr> <td style=\"text-align:center\">500</td> <td>Internal server error while processing the request.</td> </tr> </tbody> </table> <h4>API Error Codes</h4> <table> <thead> <tr> <th style=\"text-align:center\"><strong>Code</strong></th> <th><strong>Description</strong></th> </tr> </thead> <tbody> <tr> <td style=\"text-align:center\">40001</td> <td>HTTP request header validation failed.</td> </tr> <tr> <td style=\"text-align:center\">40002</td> <td>API request parameter validation failed.</td> </tr> <tr> <td style=\"text-align:center\">40101</td> <td>HTTP request header <code>Authorization</code> is not provided when required.</td> </tr> <tr> <td style=\"text-align:center\">40102</td> <td>HTTP request header <code>Authorization</code> format is invalid.</td> </tr> <tr> <td style=\"text-align:center\">40103</td> <td>Failed to parse the token or the token is not found or invalid. Client should get a new token.</td> </tr> <tr> <td style=\"text-align:center\">40104</td> <td>The token has expired. Client should refresh the access token or get a new token.</td> </tr> <tr> <td style=\"text-align:center\">40105</td> <td>The token owner does not belong to the client's info. Client should get a new token.</td> </tr> <tr> <td style=\"text-align:center\">40106</td> <td>User is not found for the specified token.</td> </tr> <tr> <td style=\"text-align:center\">40301</td> <td>The user does not have access to the requested resource or action.</td> </tr> <tr> <td style=\"text-align:center\">40401</td> <td>The requested resource is not found.</td> </tr> <tr> <td style=\"text-align:center\">49101</td> <td>The API key is not provided.</td> </tr> <tr> <td style=\"text-align:center\">49102</td> <td>Failed to parse the API key, or the API key is invalid.</td> </tr> <tr> <td style=\"text-align:center\">49103</td> <td>The provided API key is not found.</td> </tr> <tr> <td style=\"text-align:center\">49104</td> <td>The provided API key is not intended to be used with the client's app platform.</td> </tr> <tr> <td style=\"text-align:center\">49105</td> <td>The provided API key is not intended to be used with the client's app identifier.</td> </tr> <tr> <td style=\"text-align:center\">49106</td> <td>The API key has expired.</td> </tr> <tr> <td style=\"text-align:center\">49107</td> <td>The API key is disabled.</td> </tr> <tr> <td style=\"text-align:center\">50001</td> <td>An error occurred while validating the API key.</td> </tr> <tr> <td style=\"text-align:center\">99999</td> <td>Other errors, usually without specific reason or action.</td> </tr> </tbody> </table>"
  },
  {
    "type": "post",
    "url": "/v1/account/security/change_password",
    "title": "Security - Change Password",
    "version": "1.0.0",
    "name": "ChangePassword",
    "group": "AccountAPI",
    "permission": [
      {
        "name": "account"
      }
    ],
    "description": "<p>Change the account's password.</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "password",
            "description": "<p>The account's current password, to verify the request.</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "newPassword",
            "description": "<p>The new password.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Request Example:",
          "content": "{\n  \"password\": \"this_is_password\",\n  \"newPassword\": \"this_is_new_password\"\n}",
          "type": "json"
        }
      ]
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "boolean",
            "optional": false,
            "field": "success",
            "description": "<p>If the password is changed successfully.</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "message",
            "description": "<p>The message, if failed.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success Response:",
          "content": "HTTP/1.1 200 OK\n{\n  \"status\": 200,\n  \"error\": {\n    \"code\": \"\",\n    \"message\": \"\",\n    \"field\": \"\"\n  },\n  \"data\": {\n    \"success\": true,\n    \"message\": \"Password changed successfully\"\n  }\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "PasswordInvalid",
            "description": "<p>The current password is invalid.</p>"
          },
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "NewPasswordFormatInvalid",
            "description": "<p>The new password format is invalid.</p>"
          },
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "HeaderValidationFailed",
            "description": "<p>The request header validation failed.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "PasswordInvalid:",
          "content": "HTTP/1.1 200 OK\n{\n  \"status\": 400,\n  \"error\": {\n    \"code\": \"40002\",\n    \"message\": \"Current password is invalid\",\n    \"field\": \"password\"\n  },\n  \"data\": {}\n}",
          "type": "json"
        },
        {
          "title": "NewPasswordFormatInvalid:",
          "content": "HTTP/1.1 200 OK\n{\n  \"status\": 400,\n  \"error\": {\n    \"code\": \"40002\",\n    \"message\": \"Password must contain at least 1 lowercase, uppercase, and special characters and 1 number\",\n    \"field\": \"newPassword\"\n  },\n  \"data\": {}\n}",
          "type": "json"
        },
        {
          "title": "HeaderValidationFailed:",
          "content": "HTTP/1.1 200 OK\n{\n  \"status\": 400,\n  \"error\": {\n    \"code\": \"40001\",\n    \"message\": \"Request headers are required (API-Key, Authorization, Device-Platform)\",\n    \"field\": \"\"\n  },\n  \"data\": {}\n}",
          "type": "json"
        }
      ]
    },
    "filename": "internal/app/basego-api/v1/endpoint/accountapi/security-change-password.go",
    "groupTitle": "Account API",
    "groupDescription": "<h4>HTTP Request Headers</h4> <table> <thead> <tr> <th><strong>Header Name</strong></th> <th style=\"text-align:center\"><strong>Required</strong></th> <th><strong>Description</strong></th> </tr> </thead> <tbody> <tr> <td>API-Key</td> <td style=\"text-align:center\">✓</td> <td>API key for accessing the API.</td> </tr> <tr> <td>App-Identifier</td> <td style=\"text-align:center\"></td> <td>The app's identifier (package name for Android, bundle ID for iOS, or origin URL for web).</td> </tr> <tr> <td>Authorization</td> <td style=\"text-align:center\">✓</td> <td>Access token to validate the user session.<br>Format: <code>Bearer <i>&lt;access_token&gt;</i></code></td> </tr> <tr> <td>Content-Type</td> <td style=\"text-align:center\"></td> <td>Content type of the request body.</td> </tr> <tr> <td>Device-Identifier</td> <td style=\"text-align:center\">✓</td> <td>The device ID (optional for web).</td> </tr> <tr> <td>Device-Model</td> <td style=\"text-align:center\">✓</td> <td>Model name of the device (optional for web).</td> </tr> <tr> <td>Device-Platform</td> <td style=\"text-align:center\">✓</td> <td>The device's platform. Values are <code>android</code>, <code>ios</code>, or <code>web</code>.</td> </tr> <tr> <td>User-Agent</td> <td style=\"text-align:center\"></td> <td>The user agent of the client accessing the API.</td> </tr> </tbody> </table> <h4>HTTP Response Status Codes</h4> <table> <thead> <tr> <th style=\"text-align:center\"><strong>Code</strong></th> <th><strong>Description</strong></th> </tr> </thead> <tbody> <tr> <td style=\"text-align:center\">200</td> <td>OK, request proceed without error.</td> </tr> <tr> <td style=\"text-align:center\">204</td> <td>Request proceed successfully and is not returning any content (empty <code>data</code>).</td> </tr> <tr> <td style=\"text-align:center\">400</td> <td>Bad request, either the request header or the API parameter validation failed.</td> </tr> <tr> <td style=\"text-align:center\">401</td> <td>Unauthorized access, either the HTTP request header <code>Authorization</code> is missing or invalid.</td> </tr> <tr> <td style=\"text-align:center\">403</td> <td>Forbidden, the user does not have access to the requested resource.</td> </tr> <tr> <td style=\"text-align:center\">404</td> <td>Not found, requested resource does not exist.</td> </tr> <tr> <td style=\"text-align:center\">429</td> <td>Too many requests sent in a given amount of time, intended for use with rate-limiting schemes.</td> </tr> <tr> <td style=\"text-align:center\">491</td> <td>The API key specified in HTTP request header <code>API-Key</code> is invalid.</td> </tr> <tr> <td style=\"text-align:center\">500</td> <td>Internal server error while processing the request.</td> </tr> </tbody> </table> <h4>API Error Codes</h4> <table> <thead> <tr> <th style=\"text-align:center\"><strong>Code</strong></th> <th><strong>Description</strong></th> </tr> </thead> <tbody> <tr> <td style=\"text-align:center\">40001</td> <td>HTTP request header validation failed.</td> </tr> <tr> <td style=\"text-align:center\">40002</td> <td>API request parameter validation failed.</td> </tr> <tr> <td style=\"text-align:center\">40101</td> <td>HTTP request header <code>Authorization</code> is not provided when required.</td> </tr> <tr> <td style=\"text-align:center\">40102</td> <td>HTTP request header <code>Authorization</code> format is invalid.</td> </tr> <tr> <td style=\"text-align:center\">40103</td> <td>Failed to parse the token or the token is not found or invalid. Client should get a new token.</td> </tr> <tr> <td style=\"text-align:center\">40104</td> <td>The token has expired. Client should refresh the access token or get a new token.</td> </tr> <tr> <td style=\"text-align:center\">40105</td> <td>The token owner does not belong to the client's info. Client should get a new token.</td> </tr> <tr> <td style=\"text-align:center\">40106</td> <td>User is not found for the specified token.</td> </tr> <tr> <td style=\"text-align:center\">40301</td> <td>The user does not have access to the requested resource or action.</td> </tr> <tr> <td style=\"text-align:center\">40401</td> <td>The requested resource is not found.</td> </tr> <tr> <td style=\"text-align:center\">49101</td> <td>The API key is not provided.</td> </tr> <tr> <td style=\"text-align:center\">49102</td> <td>Failed to parse the API key, or the API key is invalid.</td> </tr> <tr> <td style=\"text-align:center\">49103</td> <td>The provided API key is not found.</td> </tr> <tr> <td style=\"text-align:center\">49104</td> <td>The provided API key is not intended to be used with the client's app platform.</td> </tr> <tr> <td style=\"text-align:center\">49105</td> <td>The provided API key is not intended to be used with the client's app identifier.</td> </tr> <tr> <td style=\"text-align:center\">49106</td> <td>The API key has expired.</td> </tr> <tr> <td style=\"text-align:center\">49107</td> <td>The API key is disabled.</td> </tr> <tr> <td style=\"text-align:center\">50001</td> <td>An error occurred while validating the API key.</td> </tr> <tr> <td style=\"text-align:center\">99999</td> <td>Other errors, usually without specific reason or action.</td> </tr> </tbody> </table>"
  },
  {
    "type": "post",
    "url": "/v1/account/countries",
    "title": "Get Country List",
    "version": "1.0.0",
    "name": "Countries",
    "group": "AccountAPI",
    "permission": [
      {
        "name": "account"
      }
    ],
    "description": "<p>Get the country list.</p>",
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "object[]",
            "optional": false,
            "field": "countries",
            "description": "<p>The list of countries.</p>"
          },
          {
            "group": "Success 200",
            "type": "id",
            "optional": false,
            "field": "countries.id",
            "description": "<p>The country ID.</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "countries.commonName",
            "description": "<p>The country's common name.</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "countries.officialName",
            "description": "<p>The country's official name.</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "countries.iso2Code",
            "description": "<p>The country code based on ISO 3166-1 alpha-2.</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "countries.iso3Code",
            "description": "<p>The country code based on ISO 3166-1 alpha-3.</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "countries.callingCode",
            "description": "<p>The country's international calling code for phone number.</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "countries.currencyCode",
            "description": "<p>The country's currency (e.g.: &quot;IDR&quot;).</p>"
          },
          {
            "group": "Success 200",
            "type": "boolean",
            "optional": false,
            "field": "countries.isEnabled",
            "description": "<p>If the country is enabled.</p>"
          },
          {
            "group": "Success 200",
            "type": "boolean",
            "optional": false,
            "field": "countries.isHidden",
            "description": "<p>If the country is hidden.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success Response:",
          "content": "HTTP/1.1 200 OK\n{\n  \"status\": 200,\n  \"error\": {\n    \"code\": \"\",\n    \"message\": \"\",\n    \"field\": \"\"\n  },\n  \"data\": {\n    \"countries\": [\n       {\n         \"id\": 1,\n         \"commonName\": \"Indonesia\",\n         \"officialName\": \"Republic of Indonesia\",\n         \"iso2Code\": \"ID\",\n         \"iso3Code\": \"IDN\",\n         \"callingCode\": \"62\",\n         \"currencyCode\": \"IDR\",\n         \"isEnabled\": true,\n         \"isHidden\": false\n       },\n       ...\n    ]\n  }\n}",
          "type": "json"
        }
      ]
    },
    "filename": "internal/app/basego-api/v1/endpoint/accountapi/countries.go",
    "groupTitle": "Account API",
    "groupDescription": "<h4>HTTP Request Headers</h4> <table> <thead> <tr> <th><strong>Header Name</strong></th> <th style=\"text-align:center\"><strong>Required</strong></th> <th><strong>Description</strong></th> </tr> </thead> <tbody> <tr> <td>API-Key</td> <td style=\"text-align:center\">✓</td> <td>API key for accessing the API.</td> </tr> <tr> <td>App-Identifier</td> <td style=\"text-align:center\"></td> <td>The app's identifier (package name for Android, bundle ID for iOS, or origin URL for web).</td> </tr> <tr> <td>Authorization</td> <td style=\"text-align:center\">✓</td> <td>Access token to validate the user session.<br>Format: <code>Bearer <i>&lt;access_token&gt;</i></code></td> </tr> <tr> <td>Content-Type</td> <td style=\"text-align:center\"></td> <td>Content type of the request body.</td> </tr> <tr> <td>Device-Identifier</td> <td style=\"text-align:center\">✓</td> <td>The device ID (optional for web).</td> </tr> <tr> <td>Device-Model</td> <td style=\"text-align:center\">✓</td> <td>Model name of the device (optional for web).</td> </tr> <tr> <td>Device-Platform</td> <td style=\"text-align:center\">✓</td> <td>The device's platform. Values are <code>android</code>, <code>ios</code>, or <code>web</code>.</td> </tr> <tr> <td>User-Agent</td> <td style=\"text-align:center\"></td> <td>The user agent of the client accessing the API.</td> </tr> </tbody> </table> <h4>HTTP Response Status Codes</h4> <table> <thead> <tr> <th style=\"text-align:center\"><strong>Code</strong></th> <th><strong>Description</strong></th> </tr> </thead> <tbody> <tr> <td style=\"text-align:center\">200</td> <td>OK, request proceed without error.</td> </tr> <tr> <td style=\"text-align:center\">204</td> <td>Request proceed successfully and is not returning any content (empty <code>data</code>).</td> </tr> <tr> <td style=\"text-align:center\">400</td> <td>Bad request, either the request header or the API parameter validation failed.</td> </tr> <tr> <td style=\"text-align:center\">401</td> <td>Unauthorized access, either the HTTP request header <code>Authorization</code> is missing or invalid.</td> </tr> <tr> <td style=\"text-align:center\">403</td> <td>Forbidden, the user does not have access to the requested resource.</td> </tr> <tr> <td style=\"text-align:center\">404</td> <td>Not found, requested resource does not exist.</td> </tr> <tr> <td style=\"text-align:center\">429</td> <td>Too many requests sent in a given amount of time, intended for use with rate-limiting schemes.</td> </tr> <tr> <td style=\"text-align:center\">491</td> <td>The API key specified in HTTP request header <code>API-Key</code> is invalid.</td> </tr> <tr> <td style=\"text-align:center\">500</td> <td>Internal server error while processing the request.</td> </tr> </tbody> </table> <h4>API Error Codes</h4> <table> <thead> <tr> <th style=\"text-align:center\"><strong>Code</strong></th> <th><strong>Description</strong></th> </tr> </thead> <tbody> <tr> <td style=\"text-align:center\">40001</td> <td>HTTP request header validation failed.</td> </tr> <tr> <td style=\"text-align:center\">40002</td> <td>API request parameter validation failed.</td> </tr> <tr> <td style=\"text-align:center\">40101</td> <td>HTTP request header <code>Authorization</code> is not provided when required.</td> </tr> <tr> <td style=\"text-align:center\">40102</td> <td>HTTP request header <code>Authorization</code> format is invalid.</td> </tr> <tr> <td style=\"text-align:center\">40103</td> <td>Failed to parse the token or the token is not found or invalid. Client should get a new token.</td> </tr> <tr> <td style=\"text-align:center\">40104</td> <td>The token has expired. Client should refresh the access token or get a new token.</td> </tr> <tr> <td style=\"text-align:center\">40105</td> <td>The token owner does not belong to the client's info. Client should get a new token.</td> </tr> <tr> <td style=\"text-align:center\">40106</td> <td>User is not found for the specified token.</td> </tr> <tr> <td style=\"text-align:center\">40301</td> <td>The user does not have access to the requested resource or action.</td> </tr> <tr> <td style=\"text-align:center\">40401</td> <td>The requested resource is not found.</td> </tr> <tr> <td style=\"text-align:center\">49101</td> <td>The API key is not provided.</td> </tr> <tr> <td style=\"text-align:center\">49102</td> <td>Failed to parse the API key, or the API key is invalid.</td> </tr> <tr> <td style=\"text-align:center\">49103</td> <td>The provided API key is not found.</td> </tr> <tr> <td style=\"text-align:center\">49104</td> <td>The provided API key is not intended to be used with the client's app platform.</td> </tr> <tr> <td style=\"text-align:center\">49105</td> <td>The provided API key is not intended to be used with the client's app identifier.</td> </tr> <tr> <td style=\"text-align:center\">49106</td> <td>The API key has expired.</td> </tr> <tr> <td style=\"text-align:center\">49107</td> <td>The API key is disabled.</td> </tr> <tr> <td style=\"text-align:center\">50001</td> <td>An error occurred while validating the API key.</td> </tr> <tr> <td style=\"text-align:center\">99999</td> <td>Other errors, usually without specific reason or action.</td> </tr> </tbody> </table>",
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "HeaderValidationFailed",
            "description": "<p>The request header validation failed.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "HeaderValidationFailed:",
          "content": "HTTP/1.1 200 OK\n{\n  \"status\": 400,\n  \"error\": {\n    \"code\": \"40001\",\n    \"message\": \"Request headers are required (API-Key, Authorization, Device-Platform)\",\n    \"field\": \"\"\n  },\n  \"data\": {}\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "post",
    "url": "/v1/account/profile/get",
    "title": "Get Account Profile",
    "version": "1.0.0",
    "name": "GetAccountProfile",
    "group": "AccountAPI",
    "permission": [
      {
        "name": "account"
      }
    ],
    "description": "<p>Get the current account's profile.</p>",
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "object",
            "optional": false,
            "field": "account",
            "description": "<p>The account's profile.</p>"
          },
          {
            "group": "Success 200",
            "type": "object",
            "optional": false,
            "field": "tos",
            "description": "<p>The account's Terms of Service status.</p>"
          },
          {
            "group": "Success 200",
            "type": "boolean",
            "optional": false,
            "field": "tos.isAccepted",
            "description": "<p>If the Terms of Service has been accepted.</p>"
          },
          {
            "group": "Success 200",
            "type": "long",
            "optional": false,
            "field": "tos.acceptedTime",
            "description": "<p>The time the Terms of Service was accepted, in Unix milliseconds.</p>"
          },
          {
            "group": "Success 200",
            "type": "long",
            "optional": false,
            "field": "account.id",
            "description": "<p>The account ID.</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "account.fullName",
            "description": "<p>The account's full name.</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "account.email",
            "description": "<p>The email address.</p>"
          },
          {
            "group": "Success 200",
            "type": "boolean",
            "optional": false,
            "field": "account.isEmailVerified",
            "description": "<p>If the email address is verified.</p>"
          },
          {
            "group": "Success 200",
            "type": "integer",
            "optional": false,
            "field": "account.countryID",
            "description": "<p>The account's country ID.</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "account.countryCallingCode",
            "description": "<p>The country calling code for phone number.</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "account.phone",
            "description": "<p>The phone number.</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "account.phoneWithCode",
            "description": "<p>The phone number with country calling code.</p>"
          },
          {
            "group": "Success 200",
            "type": "boolean",
            "optional": false,
            "field": "account.isPhoneVerified",
            "description": "<p>If the phone number is verified.</p>"
          },
          {
            "group": "Success 200",
            "type": "object",
            "optional": false,
            "field": "account.imageURL",
            "description": "<p>The account's picture image URL.</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "account.imageURL.thumbnail",
            "description": "<p>Image URL for thumbnail picture.</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "account.imageURL.fullsize",
            "description": "<p>Image URL for fullsize picture.</p>"
          },
          {
            "group": "Success 200",
            "type": "long",
            "optional": false,
            "field": "account.lastLoginTime",
            "description": "<p>The account's last login, in Unix milliseconds.</p>"
          },
          {
            "group": "Success 200",
            "type": "long",
            "optional": false,
            "field": "account.lastActivityTime",
            "description": "<p>The account's last activity, in Unix milliseconds.</p>"
          },
          {
            "group": "Success 200",
            "type": "boolean",
            "optional": false,
            "field": "account.requireChangePassword",
            "description": "<p>If the account is required to change password.</p>"
          },
          {
            "group": "Success 200",
            "type": "long",
            "optional": false,
            "field": "account.createdTime",
            "description": "<p>The time the account was created, in Unix milliseconds.</p>"
          },
          {
            "group": "Success 200",
            "type": "long",
            "optional": false,
            "field": "account.updatedTime",
            "description": "<p>The time the account was last updated, in Unix milliseconds.</p>"
          },
          {
            "group": "Success 200",
            "type": "long",
            "optional": false,
            "field": "account.deletedTime",
            "description": "<p>The time the account was deleted, in Unix milliseconds.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success Response:",
          "content": "HTTP/1.1 200 OK\n{\n  \"status\": 200,\n  \"error\": {\n    \"code\": \"\",\n    \"message\": \"\",\n    \"field\": \"\"\n  },\n  \"data\": {\n    \"account\": {\n      \"id\": 8,\n      \"fullName\": \"Jony\",\n      \"email\": \"\",\n      \"isEmailVerified\": false,\n      \"countryID\": 0,\n      \"countryCallingCode\": \"\",\n      \"phone\": \"\",\n      \"phoneWithCode\": \"\",\n      \"isPhoneVerified\": false,\n      \"imageURL\": {\n        \"thumbnail\": \"\",\n        \"fullsize\": \"\"\n      },\n      \"lastLoginTime\": 1564121972641,\n      \"lastActivityTime\": 1564121972641,\n      \"requireChangePassword\": false,\n      \"createdTime\": 1563868799147,\n      \"updatedTime\": 1563880378559,\n      \"deletedTime\": 0\n    },\n    \"tos\": {\n      \"isAccepted\": true,\n      \"acceptedTime\": 1566452967572\n    }\n  }\n}",
          "type": "json"
        }
      ]
    },
    "filename": "internal/app/basego-api/v1/endpoint/accountapi/profile-get.go",
    "groupTitle": "Account API",
    "groupDescription": "<h4>HTTP Request Headers</h4> <table> <thead> <tr> <th><strong>Header Name</strong></th> <th style=\"text-align:center\"><strong>Required</strong></th> <th><strong>Description</strong></th> </tr> </thead> <tbody> <tr> <td>API-Key</td> <td style=\"text-align:center\">✓</td> <td>API key for accessing the API.</td> </tr> <tr> <td>App-Identifier</td> <td style=\"text-align:center\"></td> <td>The app's identifier (package name for Android, bundle ID for iOS, or origin URL for web).</td> </tr> <tr> <td>Authorization</td> <td style=\"text-align:center\">✓</td> <td>Access token to validate the user session.<br>Format: <code>Bearer <i>&lt;access_token&gt;</i></code></td> </tr> <tr> <td>Content-Type</td> <td style=\"text-align:center\"></td> <td>Content type of the request body.</td> </tr> <tr> <td>Device-Identifier</td> <td style=\"text-align:center\">✓</td> <td>The device ID (optional for web).</td> </tr> <tr> <td>Device-Model</td> <td style=\"text-align:center\">✓</td> <td>Model name of the device (optional for web).</td> </tr> <tr> <td>Device-Platform</td> <td style=\"text-align:center\">✓</td> <td>The device's platform. Values are <code>android</code>, <code>ios</code>, or <code>web</code>.</td> </tr> <tr> <td>User-Agent</td> <td style=\"text-align:center\"></td> <td>The user agent of the client accessing the API.</td> </tr> </tbody> </table> <h4>HTTP Response Status Codes</h4> <table> <thead> <tr> <th style=\"text-align:center\"><strong>Code</strong></th> <th><strong>Description</strong></th> </tr> </thead> <tbody> <tr> <td style=\"text-align:center\">200</td> <td>OK, request proceed without error.</td> </tr> <tr> <td style=\"text-align:center\">204</td> <td>Request proceed successfully and is not returning any content (empty <code>data</code>).</td> </tr> <tr> <td style=\"text-align:center\">400</td> <td>Bad request, either the request header or the API parameter validation failed.</td> </tr> <tr> <td style=\"text-align:center\">401</td> <td>Unauthorized access, either the HTTP request header <code>Authorization</code> is missing or invalid.</td> </tr> <tr> <td style=\"text-align:center\">403</td> <td>Forbidden, the user does not have access to the requested resource.</td> </tr> <tr> <td style=\"text-align:center\">404</td> <td>Not found, requested resource does not exist.</td> </tr> <tr> <td style=\"text-align:center\">429</td> <td>Too many requests sent in a given amount of time, intended for use with rate-limiting schemes.</td> </tr> <tr> <td style=\"text-align:center\">491</td> <td>The API key specified in HTTP request header <code>API-Key</code> is invalid.</td> </tr> <tr> <td style=\"text-align:center\">500</td> <td>Internal server error while processing the request.</td> </tr> </tbody> </table> <h4>API Error Codes</h4> <table> <thead> <tr> <th style=\"text-align:center\"><strong>Code</strong></th> <th><strong>Description</strong></th> </tr> </thead> <tbody> <tr> <td style=\"text-align:center\">40001</td> <td>HTTP request header validation failed.</td> </tr> <tr> <td style=\"text-align:center\">40002</td> <td>API request parameter validation failed.</td> </tr> <tr> <td style=\"text-align:center\">40101</td> <td>HTTP request header <code>Authorization</code> is not provided when required.</td> </tr> <tr> <td style=\"text-align:center\">40102</td> <td>HTTP request header <code>Authorization</code> format is invalid.</td> </tr> <tr> <td style=\"text-align:center\">40103</td> <td>Failed to parse the token or the token is not found or invalid. Client should get a new token.</td> </tr> <tr> <td style=\"text-align:center\">40104</td> <td>The token has expired. Client should refresh the access token or get a new token.</td> </tr> <tr> <td style=\"text-align:center\">40105</td> <td>The token owner does not belong to the client's info. Client should get a new token.</td> </tr> <tr> <td style=\"text-align:center\">40106</td> <td>User is not found for the specified token.</td> </tr> <tr> <td style=\"text-align:center\">40301</td> <td>The user does not have access to the requested resource or action.</td> </tr> <tr> <td style=\"text-align:center\">40401</td> <td>The requested resource is not found.</td> </tr> <tr> <td style=\"text-align:center\">49101</td> <td>The API key is not provided.</td> </tr> <tr> <td style=\"text-align:center\">49102</td> <td>Failed to parse the API key, or the API key is invalid.</td> </tr> <tr> <td style=\"text-align:center\">49103</td> <td>The provided API key is not found.</td> </tr> <tr> <td style=\"text-align:center\">49104</td> <td>The provided API key is not intended to be used with the client's app platform.</td> </tr> <tr> <td style=\"text-align:center\">49105</td> <td>The provided API key is not intended to be used with the client's app identifier.</td> </tr> <tr> <td style=\"text-align:center\">49106</td> <td>The API key has expired.</td> </tr> <tr> <td style=\"text-align:center\">49107</td> <td>The API key is disabled.</td> </tr> <tr> <td style=\"text-align:center\">50001</td> <td>An error occurred while validating the API key.</td> </tr> <tr> <td style=\"text-align:center\">99999</td> <td>Other errors, usually without specific reason or action.</td> </tr> </tbody> </table>",
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "HeaderValidationFailed",
            "description": "<p>The request header validation failed.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "HeaderValidationFailed:",
          "content": "HTTP/1.1 200 OK\n{\n  \"status\": 400,\n  \"error\": {\n    \"code\": \"40001\",\n    \"message\": \"Request headers are required (API-Key, Authorization, Device-Platform)\",\n    \"field\": \"\"\n  },\n  \"data\": {}\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "post",
    "url": "/v1/account/logout",
    "title": "Logout",
    "version": "1.0.0",
    "name": "Logout",
    "group": "AccountAPI",
    "permission": [
      {
        "name": "account"
      }
    ],
    "description": "<p>Log out of an account session.</p>",
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "boolean",
            "optional": false,
            "field": "success",
            "description": "<p>If the logout is successful.</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "message",
            "description": "<p>The message.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success Response:",
          "content": "HTTP/1.1 200 OK\n{\n  \"status\": 200,\n  \"error\": {\n    \"code\": \"\",\n    \"message\": \"\",\n    \"field\": \"\"\n  },\n  \"data\": {\n    \"success\": true,\n    \"message\": \"You have logged out\"\n  }\n}",
          "type": "json"
        }
      ]
    },
    "filename": "internal/app/basego-api/v1/endpoint/accountapi/logout.go",
    "groupTitle": "Account API",
    "groupDescription": "<h4>HTTP Request Headers</h4> <table> <thead> <tr> <th><strong>Header Name</strong></th> <th style=\"text-align:center\"><strong>Required</strong></th> <th><strong>Description</strong></th> </tr> </thead> <tbody> <tr> <td>API-Key</td> <td style=\"text-align:center\">✓</td> <td>API key for accessing the API.</td> </tr> <tr> <td>App-Identifier</td> <td style=\"text-align:center\"></td> <td>The app's identifier (package name for Android, bundle ID for iOS, or origin URL for web).</td> </tr> <tr> <td>Authorization</td> <td style=\"text-align:center\">✓</td> <td>Access token to validate the user session.<br>Format: <code>Bearer <i>&lt;access_token&gt;</i></code></td> </tr> <tr> <td>Content-Type</td> <td style=\"text-align:center\"></td> <td>Content type of the request body.</td> </tr> <tr> <td>Device-Identifier</td> <td style=\"text-align:center\">✓</td> <td>The device ID (optional for web).</td> </tr> <tr> <td>Device-Model</td> <td style=\"text-align:center\">✓</td> <td>Model name of the device (optional for web).</td> </tr> <tr> <td>Device-Platform</td> <td style=\"text-align:center\">✓</td> <td>The device's platform. Values are <code>android</code>, <code>ios</code>, or <code>web</code>.</td> </tr> <tr> <td>User-Agent</td> <td style=\"text-align:center\"></td> <td>The user agent of the client accessing the API.</td> </tr> </tbody> </table> <h4>HTTP Response Status Codes</h4> <table> <thead> <tr> <th style=\"text-align:center\"><strong>Code</strong></th> <th><strong>Description</strong></th> </tr> </thead> <tbody> <tr> <td style=\"text-align:center\">200</td> <td>OK, request proceed without error.</td> </tr> <tr> <td style=\"text-align:center\">204</td> <td>Request proceed successfully and is not returning any content (empty <code>data</code>).</td> </tr> <tr> <td style=\"text-align:center\">400</td> <td>Bad request, either the request header or the API parameter validation failed.</td> </tr> <tr> <td style=\"text-align:center\">401</td> <td>Unauthorized access, either the HTTP request header <code>Authorization</code> is missing or invalid.</td> </tr> <tr> <td style=\"text-align:center\">403</td> <td>Forbidden, the user does not have access to the requested resource.</td> </tr> <tr> <td style=\"text-align:center\">404</td> <td>Not found, requested resource does not exist.</td> </tr> <tr> <td style=\"text-align:center\">429</td> <td>Too many requests sent in a given amount of time, intended for use with rate-limiting schemes.</td> </tr> <tr> <td style=\"text-align:center\">491</td> <td>The API key specified in HTTP request header <code>API-Key</code> is invalid.</td> </tr> <tr> <td style=\"text-align:center\">500</td> <td>Internal server error while processing the request.</td> </tr> </tbody> </table> <h4>API Error Codes</h4> <table> <thead> <tr> <th style=\"text-align:center\"><strong>Code</strong></th> <th><strong>Description</strong></th> </tr> </thead> <tbody> <tr> <td style=\"text-align:center\">40001</td> <td>HTTP request header validation failed.</td> </tr> <tr> <td style=\"text-align:center\">40002</td> <td>API request parameter validation failed.</td> </tr> <tr> <td style=\"text-align:center\">40101</td> <td>HTTP request header <code>Authorization</code> is not provided when required.</td> </tr> <tr> <td style=\"text-align:center\">40102</td> <td>HTTP request header <code>Authorization</code> format is invalid.</td> </tr> <tr> <td style=\"text-align:center\">40103</td> <td>Failed to parse the token or the token is not found or invalid. Client should get a new token.</td> </tr> <tr> <td style=\"text-align:center\">40104</td> <td>The token has expired. Client should refresh the access token or get a new token.</td> </tr> <tr> <td style=\"text-align:center\">40105</td> <td>The token owner does not belong to the client's info. Client should get a new token.</td> </tr> <tr> <td style=\"text-align:center\">40106</td> <td>User is not found for the specified token.</td> </tr> <tr> <td style=\"text-align:center\">40301</td> <td>The user does not have access to the requested resource or action.</td> </tr> <tr> <td style=\"text-align:center\">40401</td> <td>The requested resource is not found.</td> </tr> <tr> <td style=\"text-align:center\">49101</td> <td>The API key is not provided.</td> </tr> <tr> <td style=\"text-align:center\">49102</td> <td>Failed to parse the API key, or the API key is invalid.</td> </tr> <tr> <td style=\"text-align:center\">49103</td> <td>The provided API key is not found.</td> </tr> <tr> <td style=\"text-align:center\">49104</td> <td>The provided API key is not intended to be used with the client's app platform.</td> </tr> <tr> <td style=\"text-align:center\">49105</td> <td>The provided API key is not intended to be used with the client's app identifier.</td> </tr> <tr> <td style=\"text-align:center\">49106</td> <td>The API key has expired.</td> </tr> <tr> <td style=\"text-align:center\">49107</td> <td>The API key is disabled.</td> </tr> <tr> <td style=\"text-align:center\">50001</td> <td>An error occurred while validating the API key.</td> </tr> <tr> <td style=\"text-align:center\">99999</td> <td>Other errors, usually without specific reason or action.</td> </tr> </tbody> </table>",
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "HeaderValidationFailed",
            "description": "<p>The request header validation failed.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "HeaderValidationFailed:",
          "content": "HTTP/1.1 200 OK\n{\n  \"status\": 400,\n  \"error\": {\n    \"code\": \"40001\",\n    \"message\": \"Request headers are required (API-Key, Authorization, Device-Platform)\",\n    \"field\": \"\"\n  },\n  \"data\": {}\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "post",
    "url": "/v1/account/time_zones",
    "title": "Get Time Zone List",
    "version": "1.0.0",
    "name": "TimeZones",
    "group": "AccountAPI",
    "permission": [
      {
        "name": "account"
      }
    ],
    "description": "<p>Get the list of time zones.</p>",
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "object",
            "optional": false,
            "field": "timeZones",
            "description": "<p>The list of time zones.</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "timeZones.name",
            "description": "<p>The time zone's name.</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "timeZones.abbrev",
            "description": "<p>The time zone's abbrev.</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "timeZones.utcOffset",
            "description": "<p>The offset from UTC (format: <code>&quot;+HH:mm:ss&quot;</code>).</p>"
          },
          {
            "group": "Success 200",
            "type": "boolean",
            "optional": false,
            "field": "timeZones.isDST",
            "description": "<p>If the time zone is currently in DST.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success Response:",
          "content": "HTTP/1.1 200 OK\n{\n  \"status\": 200,\n  \"error\": {\n    \"code\": \"\",\n    \"message\": \"\",\n    \"field\": \"\"\n  },\n  \"data\": {\n    \"timeZones\": [\n      {\n        \"name\": \"Asia/Jakarta\",\n        \"abbrev\": \"WIB\",\n        \"utcOffset\": \"07:00:00\",\n        \"isDST\": false\n      },\n      {\n        \"name\": \"America/New_York\",\n        \"abbrev\": \"EDT\",\n        \"utcOffset\": \"-04:00:00\",\n        \"isDST\": true\n      }\n    ]\n  }\n}",
          "type": "json"
        }
      ]
    },
    "filename": "internal/app/basego-api/v1/endpoint/accountapi/time-zones.go",
    "groupTitle": "Account API",
    "groupDescription": "<h4>HTTP Request Headers</h4> <table> <thead> <tr> <th><strong>Header Name</strong></th> <th style=\"text-align:center\"><strong>Required</strong></th> <th><strong>Description</strong></th> </tr> </thead> <tbody> <tr> <td>API-Key</td> <td style=\"text-align:center\">✓</td> <td>API key for accessing the API.</td> </tr> <tr> <td>App-Identifier</td> <td style=\"text-align:center\"></td> <td>The app's identifier (package name for Android, bundle ID for iOS, or origin URL for web).</td> </tr> <tr> <td>Authorization</td> <td style=\"text-align:center\">✓</td> <td>Access token to validate the user session.<br>Format: <code>Bearer <i>&lt;access_token&gt;</i></code></td> </tr> <tr> <td>Content-Type</td> <td style=\"text-align:center\"></td> <td>Content type of the request body.</td> </tr> <tr> <td>Device-Identifier</td> <td style=\"text-align:center\">✓</td> <td>The device ID (optional for web).</td> </tr> <tr> <td>Device-Model</td> <td style=\"text-align:center\">✓</td> <td>Model name of the device (optional for web).</td> </tr> <tr> <td>Device-Platform</td> <td style=\"text-align:center\">✓</td> <td>The device's platform. Values are <code>android</code>, <code>ios</code>, or <code>web</code>.</td> </tr> <tr> <td>User-Agent</td> <td style=\"text-align:center\"></td> <td>The user agent of the client accessing the API.</td> </tr> </tbody> </table> <h4>HTTP Response Status Codes</h4> <table> <thead> <tr> <th style=\"text-align:center\"><strong>Code</strong></th> <th><strong>Description</strong></th> </tr> </thead> <tbody> <tr> <td style=\"text-align:center\">200</td> <td>OK, request proceed without error.</td> </tr> <tr> <td style=\"text-align:center\">204</td> <td>Request proceed successfully and is not returning any content (empty <code>data</code>).</td> </tr> <tr> <td style=\"text-align:center\">400</td> <td>Bad request, either the request header or the API parameter validation failed.</td> </tr> <tr> <td style=\"text-align:center\">401</td> <td>Unauthorized access, either the HTTP request header <code>Authorization</code> is missing or invalid.</td> </tr> <tr> <td style=\"text-align:center\">403</td> <td>Forbidden, the user does not have access to the requested resource.</td> </tr> <tr> <td style=\"text-align:center\">404</td> <td>Not found, requested resource does not exist.</td> </tr> <tr> <td style=\"text-align:center\">429</td> <td>Too many requests sent in a given amount of time, intended for use with rate-limiting schemes.</td> </tr> <tr> <td style=\"text-align:center\">491</td> <td>The API key specified in HTTP request header <code>API-Key</code> is invalid.</td> </tr> <tr> <td style=\"text-align:center\">500</td> <td>Internal server error while processing the request.</td> </tr> </tbody> </table> <h4>API Error Codes</h4> <table> <thead> <tr> <th style=\"text-align:center\"><strong>Code</strong></th> <th><strong>Description</strong></th> </tr> </thead> <tbody> <tr> <td style=\"text-align:center\">40001</td> <td>HTTP request header validation failed.</td> </tr> <tr> <td style=\"text-align:center\">40002</td> <td>API request parameter validation failed.</td> </tr> <tr> <td style=\"text-align:center\">40101</td> <td>HTTP request header <code>Authorization</code> is not provided when required.</td> </tr> <tr> <td style=\"text-align:center\">40102</td> <td>HTTP request header <code>Authorization</code> format is invalid.</td> </tr> <tr> <td style=\"text-align:center\">40103</td> <td>Failed to parse the token or the token is not found or invalid. Client should get a new token.</td> </tr> <tr> <td style=\"text-align:center\">40104</td> <td>The token has expired. Client should refresh the access token or get a new token.</td> </tr> <tr> <td style=\"text-align:center\">40105</td> <td>The token owner does not belong to the client's info. Client should get a new token.</td> </tr> <tr> <td style=\"text-align:center\">40106</td> <td>User is not found for the specified token.</td> </tr> <tr> <td style=\"text-align:center\">40301</td> <td>The user does not have access to the requested resource or action.</td> </tr> <tr> <td style=\"text-align:center\">40401</td> <td>The requested resource is not found.</td> </tr> <tr> <td style=\"text-align:center\">49101</td> <td>The API key is not provided.</td> </tr> <tr> <td style=\"text-align:center\">49102</td> <td>Failed to parse the API key, or the API key is invalid.</td> </tr> <tr> <td style=\"text-align:center\">49103</td> <td>The provided API key is not found.</td> </tr> <tr> <td style=\"text-align:center\">49104</td> <td>The provided API key is not intended to be used with the client's app platform.</td> </tr> <tr> <td style=\"text-align:center\">49105</td> <td>The provided API key is not intended to be used with the client's app identifier.</td> </tr> <tr> <td style=\"text-align:center\">49106</td> <td>The API key has expired.</td> </tr> <tr> <td style=\"text-align:center\">49107</td> <td>The API key is disabled.</td> </tr> <tr> <td style=\"text-align:center\">50001</td> <td>An error occurred while validating the API key.</td> </tr> <tr> <td style=\"text-align:center\">99999</td> <td>Other errors, usually without specific reason or action.</td> </tr> </tbody> </table>",
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "HeaderValidationFailed",
            "description": "<p>The request header validation failed.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "HeaderValidationFailed:",
          "content": "HTTP/1.1 200 OK\n{\n  \"status\": 400,\n  \"error\": {\n    \"code\": \"40001\",\n    \"message\": \"Request headers are required (API-Key, Authorization, Device-Platform)\",\n    \"field\": \"\"\n  },\n  \"data\": {}\n}",
          "type": "json"
        }
      ]
    }
  },
  {
    "type": "post",
    "url": "/v1/auth/access_token/refresh",
    "title": "Refresh Access Token",
    "version": "1.0.0",
    "name": "AccessToken_Refresh",
    "group": "AuthAPI",
    "permission": [
      {
        "name": "client"
      }
    ],
    "description": "<p>Request a new access token using a refresh token.</p> <p>The refresh token is passed via request header <code>Authorization</code> using the following format:</p> <pre class=\"prettyprint\">Authorization: Bearer <refresh_token> </code></pre> <p>A new refresh token will be generated and returned along with the new access token. An access token is valid for 24 hours, after which it must be refreshed using a refresh token. The old access token and refresh token will no longer be usable.</p> <p>This API has the same response structure as API <a href=\"#api-AuthAPI-AccessToken_Request\">Request Access Token</a>.</p>",
    "parameter": {
      "examples": [
        {
          "title": "Request Header Example:",
          "content": "Content-Type: application/json\nAPI-Key: YjQzYjQ4NzQ1ZGZhMGU0NGsxOk16STVYekV1TVYvOWhjSEJmYTJWNVgybGtYMkZ1WkQvb3hOVE0yT1RrM09EYzNNakkwTnpJNA==\nAuthorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0a24iOiJkMjlhNTIzZWZjNzQwNzYzNDdmM2FlOGIwOTdhNTUwM2RiZWU1YjA2ZGYxNjAwYWU2NTY4N2Y5NmFkOGNiYWRlODYzYTNlOTk0MzFiMDkxNzc0YjE1ZTNhODk1ZDkzM2Y0NDQ2NzNkNmJlNWRhZDA5M2EwYjAyMWMyMWNiNTdhNSIsInRpZCI6Mywic2lkIjozLCJ1aWQiOjgsImV4cCI6MTU2NjcxMzk3MiwianRpIjoiMTU2NDEyMTk3MnIzIiwiaWF0IjoxNTY0MTIxOTcyLCJpc3MiOiJjb25zb2xlIn0.ip_Uf2rMfgrwUB6oLJ2TVI-NT-9IpxxibdUittg3CKU\nDevice-Identifier: aaaa-bbbb-cccc-dddd\nDevice-Model: Google Chrome\nDevice-Platform: web\nUser-Agent: Google Chrome/12.1.14",
          "type": "json"
        }
      ]
    },
    "success": {
      "examples": [
        {
          "title": "Success Response:",
          "content": "HTTP/1.1 200 OK\n{\n  \"status\": 200,\n  \"error\": {\n    \"code\": \"\",\n    \"message\": \"\",\n    \"field\": \"\"\n  },\n  \"data\": {\n    \"accessToken\": \"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0a24iOiI3OWFlMmIzZmVlZGYyOTM1ZDNjODdkYWE5OGE0YjJjNTU1NmUyNjBjZDAxNjk3M2EyMjczMDA1MjhlZGQ4MGQyOTgxNWYzNWRhODk1ZTNiZDk3NTBlYWE5NTk0NGVlNWU0OGMwZGQ0NzE2MzY5OTkyYmM5MjE5Y2UzM2ExNmQwOCIsInRpZCI6Nywic2lkIjo0LCJ1aWQiOjgsImV4cCI6MTU2NDMwNTA1NiwianRpIjoiMTU2NDIxODY1NmE3IiwiaWF0IjoxNTY0MjE4NjU2LCJpc3MiOiJjb25zb2xlIn0.lLEwnoIDuAzGASvnLlZE581ljShXYPbWxtl6If1NXlc\",\n    \"accessTokenExpiry\": 1564305056000,\n    \"refreshToken\": \"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0a24iOiI4MjJmMDAxMDhkYzE4NmM1NzI3ZWI2OGExZTY5NGE2MWUxMTc2MTBkZDkzYTg0NDUwNWI5NDg4MmU2OGE5YzljMmVlNWFiMDRiODBmZDZkYjY4ZDUzMzVlNmFhNjJjM2UxNzA4NjViOTEwMzE4ODU1MTdhZmM5OGQ0YmJiZDg2YyIsInRpZCI6Nywic2lkIjo0LCJ1aWQiOjgsImV4cCI6MTU2NjgxMDY1NiwianRpIjoiMTU2NDIxODY1NnI3IiwiaWF0IjoxNTY0MjE4NjU2LCJpc3MiOiJjb25zb2xlIn0.tn5HHwENMmNUj4KwLXW3KPkMFljfsvBaTmy82vsMoeY\",\n    \"refreshTokenExpiry\": 1566810656000,\n    \"account\": {\n      \"id\": 8,\n      \"fullName\": \"Jony\",\n      \"email\": \"jony@example.com\",\n      \"isEmailVerified\": true,\n      \"countryID\": 0,\n      \"countryCallingCode\": \"\",\n      \"phone\": \"\",\n      \"phoneWithCode\": \"\",\n      \"isPhoneVerified\": false,\n      \"imageURL\": {\n        \"thumbnail\": \"\",\n        \"fullsize\": \"\"\n      },\n      \"lastLoginTime\": 1564121972641,\n      \"lastActivityTime\": 1564121972641,\n      \"requireChangePassword\": false,\n      \"createdTime\": 1563868799147,\n      \"updatedTime\": 1563880378559,\n      \"deletedTime\": 0\n    }\n  }\n}",
          "type": "json"
        }
      ],
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "accessToken",
            "description": "<p>The access token.</p>"
          },
          {
            "group": "Success 200",
            "type": "long",
            "optional": false,
            "field": "accessTokenExpiry",
            "description": "<p>The access token's expiry time, in Unix milliseconds.</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "refreshToken",
            "description": "<p>The refresh token.</p>"
          },
          {
            "group": "Success 200",
            "type": "long",
            "optional": false,
            "field": "refreshTokenExpiry",
            "description": "<p>The refresh token's expiry time, in Unix milliseconds.</p>"
          },
          {
            "group": "Success 200",
            "type": "object",
            "optional": false,
            "field": "account",
            "description": "<p>The account's profile.</p>"
          },
          {
            "group": "Success 200",
            "type": "long",
            "optional": false,
            "field": "account.id",
            "description": "<p>The account ID.</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "account.fullName",
            "description": "<p>The account's full name.</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "account.email",
            "description": "<p>The email address.</p>"
          },
          {
            "group": "Success 200",
            "type": "boolean",
            "optional": false,
            "field": "account.isEmailVerified",
            "description": "<p>If the email address is verified.</p>"
          },
          {
            "group": "Success 200",
            "type": "integer",
            "optional": false,
            "field": "account.countryID",
            "description": "<p>The account's country ID.</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "account.countryCallingCode",
            "description": "<p>The country calling code for phone number.</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "account.phone",
            "description": "<p>The phone number.</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "account.phoneWithCode",
            "description": "<p>The phone number with country calling code.</p>"
          },
          {
            "group": "Success 200",
            "type": "boolean",
            "optional": false,
            "field": "account.isPhoneVerified",
            "description": "<p>If the phone number is verified.</p>"
          },
          {
            "group": "Success 200",
            "type": "object",
            "optional": false,
            "field": "account.imageURL",
            "description": "<p>The account's picture image URL.</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "account.imageURL.thumbnail",
            "description": "<p>Image URL for thumbnail picture.</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "account.imageURL.fullsize",
            "description": "<p>Image URL for fullsize picture.</p>"
          },
          {
            "group": "Success 200",
            "type": "long",
            "optional": false,
            "field": "account.lastLoginTime",
            "description": "<p>The account's last login, in Unix milliseconds.</p>"
          },
          {
            "group": "Success 200",
            "type": "long",
            "optional": false,
            "field": "account.lastActivityTime",
            "description": "<p>The account's last activity, in Unix milliseconds.</p>"
          },
          {
            "group": "Success 200",
            "type": "boolean",
            "optional": false,
            "field": "account.requireChangePassword",
            "description": "<p>If the account is required to change password.</p>"
          },
          {
            "group": "Success 200",
            "type": "long",
            "optional": false,
            "field": "account.createdTime",
            "description": "<p>The time the account was created, in Unix milliseconds.</p>"
          },
          {
            "group": "Success 200",
            "type": "long",
            "optional": false,
            "field": "account.updatedTime",
            "description": "<p>The time the account was last updated, in Unix milliseconds.</p>"
          },
          {
            "group": "Success 200",
            "type": "long",
            "optional": false,
            "field": "account.deletedTime",
            "description": "<p>The time the account was deleted, in Unix milliseconds.</p>"
          }
        ]
      }
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "AuthorizationFormatInvalid",
            "description": "<p>The header <code>Authorization</code> format is invalid.</p>"
          },
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "RefreshTokenInvalid",
            "description": "<p>The refresh token is invalid, or has already been used.</p>"
          },
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "RefreshTokenExpired",
            "description": "<p>The refresh token has expired.</p>"
          },
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "DeviceInvalid",
            "description": "<p>The refresh token does not belong to the device.</p>"
          },
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "AccountNotFound",
            "description": "<p>The refresh token is valid, but the associated account is not found.</p>"
          },
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "HeaderValidationFailed",
            "description": "<p>The request header validation failed.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "AuthorizationFormatInvalid:",
          "content": "HTTP/1.1 200 OK\n{\n  \"status\": 401,\n  \"error\": {\n    \"code\": \"40102\",\n    \"message\": \"Authorization type is invalid or not supported\",\n    \"field\": \"\"\n  },\n  \"data\": {}\n}",
          "type": "json"
        },
        {
          "title": "RefreshTokenInvalid:",
          "content": "HTTP/1.1 200 OK\n{\n  \"status\": 401,\n  \"error\": {\n    \"code\": \"40103\",\n    \"message\": \"Token is not found\",\n    \"field\": \"\"\n  },\n  \"data\": {}\n}",
          "type": "json"
        },
        {
          "title": "RefreshTokenExpired:",
          "content": "HTTP/1.1 200 OK\n{\n  \"status\": 401,\n  \"error\": {\n    \"code\": \"40104\",\n    \"message\": \"Refresh token is expired\",\n    \"field\": \"\"\n  },\n  \"data\": {}\n}",
          "type": "json"
        },
        {
          "title": "DeviceInvalid:",
          "content": "HTTP/1.1 200 OK\n{\n  \"status\": 401,\n  \"error\": {\n    \"code\": \"40105\",\n    \"message\": \"Refresh token does not belong to the device\",\n    \"field\": \"\"\n  },\n  \"data\": {}\n}",
          "type": "json"
        },
        {
          "title": "HeaderValidationFailed:",
          "content": "HTTP/1.1 200 OK\n{\n  \"status\": 400,\n  \"error\": {\n    \"code\": \"40001\",\n    \"message\": \"Request headers are required (API-Key, Device-Platform)\",\n    \"field\": \"\"\n  },\n  \"data\": {}\n}",
          "type": "json"
        }
      ]
    },
    "filename": "internal/app/basego-api/v1/endpoint/authapi/access-token-refresh.go",
    "groupTitle": "Authentication API",
    "groupDescription": "<h4>HTTP Request Headers</h4> <table> <thead> <tr> <th><strong>Header Name</strong></th> <th style=\"text-align:center\"><strong>Required</strong></th> <th><strong>Description</strong></th> </tr> </thead> <tbody> <tr> <td>API-Key</td> <td style=\"text-align:center\">✓</td> <td>API key for accessing the API.</td> </tr> <tr> <td>App-Identifier</td> <td style=\"text-align:center\"></td> <td>The app's identifier (package name for Android, bundle ID for iOS, or origin URL for web).</td> </tr> <tr> <td>Authorization</td> <td style=\"text-align:center\">✓</td> <td>Authorization type and credentials, e.g.: basic credentials or refresh token to request new access token.<br>Format: <code><i>&lt;type&gt; &lt;credentials&gt;</i></code></td> </tr> <tr> <td>Content-Type</td> <td style=\"text-align:center\"></td> <td>Content type of the request body.</td> </tr> <tr> <td>Device-Identifier</td> <td style=\"text-align:center\">✓</td> <td>The device ID (optional for web).</td> </tr> <tr> <td>Device-Model</td> <td style=\"text-align:center\">✓</td> <td>Model name of the device (optional for web).</td> </tr> <tr> <td>Device-Platform</td> <td style=\"text-align:center\">✓</td> <td>The device's platform. Values are <code>android</code>, <code>ios</code>, or <code>web</code>.</td> </tr> <tr> <td>User-Agent</td> <td style=\"text-align:center\"></td> <td>The user agent of the client accessing the API.</td> </tr> </tbody> </table> <h4>HTTP Response Status Codes</h4> <table> <thead> <tr> <th style=\"text-align:center\"><strong>Code</strong></th> <th><strong>Description</strong></th> </tr> </thead> <tbody> <tr> <td style=\"text-align:center\">200</td> <td>OK, request proceed without error.</td> </tr> <tr> <td style=\"text-align:center\">204</td> <td>Request proceed successfully and is not returning any content (empty <code>data</code>).</td> </tr> <tr> <td style=\"text-align:center\">400</td> <td>Bad request, either the request header or the API parameter validation failed.</td> </tr> <tr> <td style=\"text-align:center\">401</td> <td>Unauthorized access, either the HTTP request header <code>Authorization</code> is missing or invalid.</td> </tr> <tr> <td style=\"text-align:center\">403</td> <td>Forbidden, the user does not have access to the requested resource.</td> </tr> <tr> <td style=\"text-align:center\">404</td> <td>Not found, requested resource does not exist.</td> </tr> <tr> <td style=\"text-align:center\">429</td> <td>Too many requests sent in a given amount of time, intended for use with rate-limiting schemes.</td> </tr> <tr> <td style=\"text-align:center\">491</td> <td>The API key specified in HTTP request header <code>API-Key</code> is invalid.</td> </tr> <tr> <td style=\"text-align:center\">500</td> <td>Internal server error while processing the request.</td> </tr> </tbody> </table> <h4>API Error Codes</h4> <table> <thead> <tr> <th style=\"text-align:center\"><strong>Code</strong></th> <th><strong>Description</strong></th> </tr> </thead> <tbody> <tr> <td style=\"text-align:center\">40001</td> <td>HTTP request header validation failed.</td> </tr> <tr> <td style=\"text-align:center\">40002</td> <td>API request parameter validation failed.</td> </tr> <tr> <td style=\"text-align:center\">40101</td> <td>HTTP request header <code>Authorization</code> is not provided when required.</td> </tr> <tr> <td style=\"text-align:center\">40102</td> <td>HTTP request header <code>Authorization</code> format is invalid.</td> </tr> <tr> <td style=\"text-align:center\">40103</td> <td>Failed to parse the token or the token is not found or invalid. Client should get a new token.</td> </tr> <tr> <td style=\"text-align:center\">40104</td> <td>The token has expired. Client should get a new token.</td> </tr> <tr> <td style=\"text-align:center\">40105</td> <td>The token owner does not belong to the client's info. Client should get a new token.</td> </tr> <tr> <td style=\"text-align:center\">40106</td> <td>User is not found for the specified token.</td> </tr> <tr> <td style=\"text-align:center\">40107</td> <td>The user's account has not been verified.</td> </tr> <tr> <td style=\"text-align:center\">40301</td> <td>The user does not have access to the requested resource or action.</td> </tr> <tr> <td style=\"text-align:center\">40401</td> <td>The requested resource is not found.</td> </tr> <tr> <td style=\"text-align:center\">49101</td> <td>The API key is not provided.</td> </tr> <tr> <td style=\"text-align:center\">49102</td> <td>Failed to parse the API key, or the API key is invalid.</td> </tr> <tr> <td style=\"text-align:center\">49103</td> <td>The provided API key is not found.</td> </tr> <tr> <td style=\"text-align:center\">49104</td> <td>The provided API key is not intended to be used with the client's app platform.</td> </tr> <tr> <td style=\"text-align:center\">49105</td> <td>The provided API key is not intended to be used with the client's app identifier.</td> </tr> <tr> <td style=\"text-align:center\">49106</td> <td>The API key has expired.</td> </tr> <tr> <td style=\"text-align:center\">49107</td> <td>The API key is disabled.</td> </tr> <tr> <td style=\"text-align:center\">50001</td> <td>An error occurred while validating the API key.</td> </tr> <tr> <td style=\"text-align:center\">99999</td> <td>Other errors, usually without specific reason or action.</td> </tr> </tbody> </table>"
  },
  {
    "type": "post",
    "url": "/v1/auth/access_token/request",
    "title": "Request Access Token",
    "version": "1.0.0",
    "name": "AccessToken_Request",
    "group": "AuthAPI",
    "permission": [
      {
        "name": "client"
      }
    ],
    "description": "<p>Request access token using the provided credentials.</p> <p>The credential is passed via request header <code>Authorization</code> using the following format:</p> <pre class=\"prettyprint\">Authorization: <type> <credentials> </code></pre> <p>The following authorization types are supported:</p> <table> <thead> <tr> <th>type</th> <th>credentials</th> </tr> </thead> <tbody> <tr> <td>Basic</td> <td><code>base64(email:password)</code></td> </tr> </tbody> </table> <p>The user session starts from the time the access token is generated. An access token is valid for 24 hours, after which it must be refreshed using a refresh token.</p> <blockquote> <p><strong>Note</strong></p> <p>Each user can only have 1 active session per device (defined by header <code>Device-Identifier</code>).<br></p> </blockquote>",
    "parameter": {
      "examples": [
        {
          "title": "Request Header Example:",
          "content": "Content-Type: application/json\nAPI-Key: YjQzYjQ4NzQ1ZGZhMGU0NGsxOk16STVYekV1TVYvOWhjSEJmYTJWNVgybGtYMkZ1WkQvb3hOVE0yT1RrM09EYzNNakkwTnpJNA==\nAuthorization: Basic am9obkBkb2UuY29tOnNlY3JldA==\nDevice-Identifier: aaaa-bbbb-cccc-dddd\nDevice-Model: Google Chrome\nDevice-Platform: web\nUser-Agent: Google Chrome/12.1.14",
          "type": "json"
        }
      ]
    },
    "success": {
      "examples": [
        {
          "title": "Success Response:",
          "content": "HTTP/1.1 200 OK\n{\n  \"status\": 200,\n  \"error\": {\n    \"code\": \"\",\n    \"message\": \"\",\n    \"field\": \"\"\n  },\n  \"data\": {\n    \"accessToken\": \"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0a24iOiIxYWM2NzI3ZjVjMGY2YTUwNTBhZWYzOWE2NDM5ZWMzOTllMWI0M2I1YTMyZGUwM2FhNjU2MGE2NTczMWM1ZDgzZmZiMzEzMjM5MjkxY2FmZDRiZjAzZGJhNjY1NTUwNzk0MzE0MWMxYjNhOTk3OTRlMjgxYzA4NTZmZjNhYjc4OSIsInRpZCI6Mywic2lkIjozLCJ1aWQiOjgsImV4cCI6MTU2NDIwODM3MiwianRpIjoiMTU2NDEyMTk3MmEzIiwiaWF0IjoxNTY0MTIxOTcyLCJpc3MiOiJjb25zb2xlIn0.ENmon7QaarCoPP3cP74kpwWNSkyXBV486VSTLwrmNCo\",\n    \"accessTokenExpiry\": 1564208372000,\n    \"refreshToken\": \"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0a24iOiJkMjlhNTIzZWZjNzQwNzYzNDdmM2FlOGIwOTdhNTUwM2RiZWU1YjA2ZGYxNjAwYWU2NTY4N2Y5NmFkOGNiYWRlODYzYTNlOTk0MzFiMDkxNzc0YjE1ZTNhODk1ZDkzM2Y0NDQ2NzNkNmJlNWRhZDA5M2EwYjAyMWMyMWNiNTdhNSIsInRpZCI6Mywic2lkIjozLCJ1aWQiOjgsImV4cCI6MTU2NjcxMzk3MiwianRpIjoiMTU2NDEyMTk3MnIzIiwiaWF0IjoxNTY0MTIxOTcyLCJpc3MiOiJjb25zb2xlIn0.ip_Uf2rMfgrwUB6oLJ2TVI-NT-9IpxxibdUittg3CKU\",\n    \"refreshTokenExpiry\": 1566713972000,\n    \"account\": {\n      \"id\": 8,\n      \"fullName\": \"Jony\",\n      \"email\": \"jony@example.com\",\n      \"isEmailVerified\": true,\n      \"countryID\": 0,\n      \"countryCallingCode\": \"\",\n      \"phone\": \"\",\n      \"phoneWithCode\": \"\",\n      \"isPhoneVerified\": false,\n      \"imageURL\": {\n        \"thumbnail\": \"\",\n        \"fullsize\": \"\"\n      },\n      \"lastLoginTime\": 1564121972641,\n      \"lastActivityTime\": 1564121972641,\n      \"requireChangePassword\": false,\n      \"createdTime\": 1563868799147,\n      \"updatedTime\": 1563880378559,\n      \"deletedTime\": 0\n    }\n  }\n}",
          "type": "json"
        }
      ],
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "accessToken",
            "description": "<p>The access token.</p>"
          },
          {
            "group": "Success 200",
            "type": "long",
            "optional": false,
            "field": "accessTokenExpiry",
            "description": "<p>The access token's expiry time, in Unix milliseconds.</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "refreshToken",
            "description": "<p>The refresh token.</p>"
          },
          {
            "group": "Success 200",
            "type": "long",
            "optional": false,
            "field": "refreshTokenExpiry",
            "description": "<p>The refresh token's expiry time, in Unix milliseconds.</p>"
          },
          {
            "group": "Success 200",
            "type": "object",
            "optional": false,
            "field": "account",
            "description": "<p>The account's profile.</p>"
          },
          {
            "group": "Success 200",
            "type": "long",
            "optional": false,
            "field": "account.id",
            "description": "<p>The account ID.</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "account.fullName",
            "description": "<p>The account's full name.</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "account.email",
            "description": "<p>The email address.</p>"
          },
          {
            "group": "Success 200",
            "type": "boolean",
            "optional": false,
            "field": "account.isEmailVerified",
            "description": "<p>If the email address is verified.</p>"
          },
          {
            "group": "Success 200",
            "type": "integer",
            "optional": false,
            "field": "account.countryID",
            "description": "<p>The account's country ID.</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "account.countryCallingCode",
            "description": "<p>The country calling code for phone number.</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "account.phone",
            "description": "<p>The phone number.</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "account.phoneWithCode",
            "description": "<p>The phone number with country calling code.</p>"
          },
          {
            "group": "Success 200",
            "type": "boolean",
            "optional": false,
            "field": "account.isPhoneVerified",
            "description": "<p>If the phone number is verified.</p>"
          },
          {
            "group": "Success 200",
            "type": "object",
            "optional": false,
            "field": "account.imageURL",
            "description": "<p>The account's picture image URL.</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "account.imageURL.thumbnail",
            "description": "<p>Image URL for thumbnail picture.</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "account.imageURL.fullsize",
            "description": "<p>Image URL for fullsize picture.</p>"
          },
          {
            "group": "Success 200",
            "type": "long",
            "optional": false,
            "field": "account.lastLoginTime",
            "description": "<p>The account's last login, in Unix milliseconds.</p>"
          },
          {
            "group": "Success 200",
            "type": "long",
            "optional": false,
            "field": "account.lastActivityTime",
            "description": "<p>The account's last activity, in Unix milliseconds.</p>"
          },
          {
            "group": "Success 200",
            "type": "boolean",
            "optional": false,
            "field": "account.requireChangePassword",
            "description": "<p>If the account is required to change password.</p>"
          },
          {
            "group": "Success 200",
            "type": "long",
            "optional": false,
            "field": "account.createdTime",
            "description": "<p>The time the account was created, in Unix milliseconds.</p>"
          },
          {
            "group": "Success 200",
            "type": "long",
            "optional": false,
            "field": "account.updatedTime",
            "description": "<p>The time the account was last updated, in Unix milliseconds.</p>"
          },
          {
            "group": "Success 200",
            "type": "long",
            "optional": false,
            "field": "account.deletedTime",
            "description": "<p>The time the account was deleted, in Unix milliseconds.</p>"
          }
        ]
      }
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "AuthorizationFormatInvalid",
            "description": "<p>The header <code>Authorization</code> format is invalid.</p>"
          },
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "CredentialsInvalid",
            "description": "<p>The credentials is invalid.</p>"
          },
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "AccountNotFound",
            "description": "<p>The account is not found.</p>"
          },
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "AccountNotVerified",
            "description": "<p>The account is not verified yet.</p>"
          },
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "HeaderValidationFailed",
            "description": "<p>The request header validation failed.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "AuthorizationFormatInvalid:",
          "content": "HTTP/1.1 200 OK\n{\n  \"status\": 401,\n  \"error\": {\n    \"code\": \"40102\",\n    \"message\": \"Authorization type is invalid or not supported\",\n    \"field\": \"\"\n  },\n  \"data\": {}\n}",
          "type": "json"
        },
        {
          "title": "CredentialsInvalid:",
          "content": "HTTP/1.1 200 OK\n{\n  \"status\": 401,\n  \"error\": {\n    \"code\": \"40103\",\n    \"message\": \"The email and password does not match\",\n    \"field\": \"\"\n  },\n  \"data\": {}\n}",
          "type": "json"
        },
        {
          "title": "AccountNotFound:",
          "content": "HTTP/1.1 200 OK\n{\n  \"status\": 401,\n  \"error\": {\n    \"code\": \"40106\",\n    \"message\": \"Account is not found\",\n    \"field\": \"\"\n  },\n  \"data\": {}\n}",
          "type": "json"
        },
        {
          "title": "AccountNotVerified:",
          "content": "HTTP/1.1 200 OK\n{\n  \"status\": 401,\n  \"error\": {\n    \"code\": \"40107\",\n    \"message\": \"Your account has not been verified yet\",\n    \"field\": \"\"\n  },\n  \"data\": {}\n}",
          "type": "json"
        },
        {
          "title": "HeaderValidationFailed:",
          "content": "HTTP/1.1 200 OK\n{\n  \"status\": 400,\n  \"error\": {\n    \"code\": \"40001\",\n    \"message\": \"Request headers are required (API-Key, Authorization, Device-Platform)\",\n    \"field\": \"\"\n  },\n  \"data\": {}\n}",
          "type": "json"
        }
      ]
    },
    "filename": "internal/app/basego-api/v1/endpoint/authapi/access-token-request.go",
    "groupTitle": "Authentication API",
    "groupDescription": "<h4>HTTP Request Headers</h4> <table> <thead> <tr> <th><strong>Header Name</strong></th> <th style=\"text-align:center\"><strong>Required</strong></th> <th><strong>Description</strong></th> </tr> </thead> <tbody> <tr> <td>API-Key</td> <td style=\"text-align:center\">✓</td> <td>API key for accessing the API.</td> </tr> <tr> <td>App-Identifier</td> <td style=\"text-align:center\"></td> <td>The app's identifier (package name for Android, bundle ID for iOS, or origin URL for web).</td> </tr> <tr> <td>Authorization</td> <td style=\"text-align:center\">✓</td> <td>Authorization type and credentials, e.g.: basic credentials or refresh token to request new access token.<br>Format: <code><i>&lt;type&gt; &lt;credentials&gt;</i></code></td> </tr> <tr> <td>Content-Type</td> <td style=\"text-align:center\"></td> <td>Content type of the request body.</td> </tr> <tr> <td>Device-Identifier</td> <td style=\"text-align:center\">✓</td> <td>The device ID (optional for web).</td> </tr> <tr> <td>Device-Model</td> <td style=\"text-align:center\">✓</td> <td>Model name of the device (optional for web).</td> </tr> <tr> <td>Device-Platform</td> <td style=\"text-align:center\">✓</td> <td>The device's platform. Values are <code>android</code>, <code>ios</code>, or <code>web</code>.</td> </tr> <tr> <td>User-Agent</td> <td style=\"text-align:center\"></td> <td>The user agent of the client accessing the API.</td> </tr> </tbody> </table> <h4>HTTP Response Status Codes</h4> <table> <thead> <tr> <th style=\"text-align:center\"><strong>Code</strong></th> <th><strong>Description</strong></th> </tr> </thead> <tbody> <tr> <td style=\"text-align:center\">200</td> <td>OK, request proceed without error.</td> </tr> <tr> <td style=\"text-align:center\">204</td> <td>Request proceed successfully and is not returning any content (empty <code>data</code>).</td> </tr> <tr> <td style=\"text-align:center\">400</td> <td>Bad request, either the request header or the API parameter validation failed.</td> </tr> <tr> <td style=\"text-align:center\">401</td> <td>Unauthorized access, either the HTTP request header <code>Authorization</code> is missing or invalid.</td> </tr> <tr> <td style=\"text-align:center\">403</td> <td>Forbidden, the user does not have access to the requested resource.</td> </tr> <tr> <td style=\"text-align:center\">404</td> <td>Not found, requested resource does not exist.</td> </tr> <tr> <td style=\"text-align:center\">429</td> <td>Too many requests sent in a given amount of time, intended for use with rate-limiting schemes.</td> </tr> <tr> <td style=\"text-align:center\">491</td> <td>The API key specified in HTTP request header <code>API-Key</code> is invalid.</td> </tr> <tr> <td style=\"text-align:center\">500</td> <td>Internal server error while processing the request.</td> </tr> </tbody> </table> <h4>API Error Codes</h4> <table> <thead> <tr> <th style=\"text-align:center\"><strong>Code</strong></th> <th><strong>Description</strong></th> </tr> </thead> <tbody> <tr> <td style=\"text-align:center\">40001</td> <td>HTTP request header validation failed.</td> </tr> <tr> <td style=\"text-align:center\">40002</td> <td>API request parameter validation failed.</td> </tr> <tr> <td style=\"text-align:center\">40101</td> <td>HTTP request header <code>Authorization</code> is not provided when required.</td> </tr> <tr> <td style=\"text-align:center\">40102</td> <td>HTTP request header <code>Authorization</code> format is invalid.</td> </tr> <tr> <td style=\"text-align:center\">40103</td> <td>Failed to parse the token or the token is not found or invalid. Client should get a new token.</td> </tr> <tr> <td style=\"text-align:center\">40104</td> <td>The token has expired. Client should get a new token.</td> </tr> <tr> <td style=\"text-align:center\">40105</td> <td>The token owner does not belong to the client's info. Client should get a new token.</td> </tr> <tr> <td style=\"text-align:center\">40106</td> <td>User is not found for the specified token.</td> </tr> <tr> <td style=\"text-align:center\">40107</td> <td>The user's account has not been verified.</td> </tr> <tr> <td style=\"text-align:center\">40301</td> <td>The user does not have access to the requested resource or action.</td> </tr> <tr> <td style=\"text-align:center\">40401</td> <td>The requested resource is not found.</td> </tr> <tr> <td style=\"text-align:center\">49101</td> <td>The API key is not provided.</td> </tr> <tr> <td style=\"text-align:center\">49102</td> <td>Failed to parse the API key, or the API key is invalid.</td> </tr> <tr> <td style=\"text-align:center\">49103</td> <td>The provided API key is not found.</td> </tr> <tr> <td style=\"text-align:center\">49104</td> <td>The provided API key is not intended to be used with the client's app platform.</td> </tr> <tr> <td style=\"text-align:center\">49105</td> <td>The provided API key is not intended to be used with the client's app identifier.</td> </tr> <tr> <td style=\"text-align:center\">49106</td> <td>The API key has expired.</td> </tr> <tr> <td style=\"text-align:center\">49107</td> <td>The API key is disabled.</td> </tr> <tr> <td style=\"text-align:center\">50001</td> <td>An error occurred while validating the API key.</td> </tr> <tr> <td style=\"text-align:center\">99999</td> <td>Other errors, usually without specific reason or action.</td> </tr> </tbody> </table>"
  },
  {
    "type": "post",
    "url": "/v1/client/account_verification/resend_email",
    "title": "Account Verification - Resend Email",
    "version": "1.0.0",
    "name": "AccountVerification_ResendEmail",
    "group": "ClientAPI",
    "permission": [
      {
        "name": "client"
      }
    ],
    "description": "<p>Resend email for account verification.</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "email",
            "description": "<p>The account's email address.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Request Example:",
          "content": "{\n  \"email\": \"john@doe.com\"\n}",
          "type": "json"
        }
      ]
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "boolean",
            "optional": false,
            "field": "success",
            "description": "<p>If the email is resent successfully.</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "message",
            "description": "<p>The message, if failed.</p>"
          },
          {
            "group": "Success 200",
            "type": "long",
            "optional": false,
            "field": "otpID",
            "description": "<p>The OTP ID for OTP verification.</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "otpKey",
            "description": "<p>The OTP key.</p>"
          },
          {
            "group": "Success 200",
            "type": "integer",
            "optional": false,
            "field": "codeLength",
            "description": "<p>The OTP code's length.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success Response:",
          "content": "HTTP/1.1 200 OK\n{\n  \"status\": 200,\n  \"error\": {\n    \"code\": \"\",\n    \"message\": \"\",\n    \"field\": \"\"\n  },\n  \"data\": {\n    \"success\": true,\n    \"message\": \"\",\n    \"otpID\": 128,\n    \"otpKey\": \"49ff390684515734c4c645df3884ed7c\",\n    \"codeLength\": 6\n  }\n}",
          "type": "json"
        },
        {
          "title": "Success Response (Already Verified):",
          "content": "HTTP/1.1 200 OK\n{\n  \"status\": 200,\n  \"error\": {\n    \"code\": \"\",\n    \"message\": \"\",\n    \"field\": \"\"\n  },\n  \"data\": {\n    \"success\": false,\n    \"message\": \"The account has already been verified\"\n  }\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "ParamValidationFailed",
            "description": "<p>The parameter validation failed.</p>"
          },
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "EmailNotRegistered",
            "description": "<p>The email address is not registered.</p>"
          },
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "HeaderValidationFailed",
            "description": "<p>The request header validation failed.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "ParamValidationFailed:",
          "content": "HTTP/1.1 200 OK\n{\n  \"status\": 400,\n  \"error\": {\n    \"code\": \"40002\",\n    \"message\": \"Email address format is invalid\",\n    \"field\": \"email\"\n  },\n  \"data\": {}\n}",
          "type": "json"
        },
        {
          "title": "EmailNotRegistered:",
          "content": "HTTP/1.1 200 OK\n{\n  \"status\": 400,\n  \"error\": {\n    \"code\": \"40002\",\n    \"message\": \"The email address is not registered\",\n    \"field\": \"email\"\n  },\n  \"data\": {}\n}",
          "type": "json"
        },
        {
          "title": "HeaderValidationFailed:",
          "content": "HTTP/1.1 200 OK\n{\n  \"status\": 400,\n  \"error\": {\n    \"code\": \"40001\",\n    \"message\": \"Request headers are required (API-Key, Device-Platform)\",\n    \"field\": \"\"\n  },\n  \"data\": {}\n}",
          "type": "json"
        }
      ]
    },
    "filename": "internal/app/basego-api/v1/endpoint/clientapi/account-verification-resend-email.go",
    "groupTitle": "Client API",
    "groupDescription": "<h4>HTTP Request Headers</h4> <table> <thead> <tr> <th><strong>Header Name</strong></th> <th style=\"text-align:center\"><strong>Required</strong></th> <th><strong>Description</strong></th> </tr> </thead> <tbody> <tr> <td>API-Key</td> <td style=\"text-align:center\">✓</td> <td>API key for accessing the API.</td> </tr> <tr> <td>App-Identifier</td> <td style=\"text-align:center\"></td> <td>The app's identifier (package name for Android, bundle ID for iOS, or origin URL for web).</td> </tr> <tr> <td>Content-Type</td> <td style=\"text-align:center\"></td> <td>Content type of the request body.</td> </tr> <tr> <td>Device-Identifier</td> <td style=\"text-align:center\">✓</td> <td>The device ID (optional for web).</td> </tr> <tr> <td>Device-Model</td> <td style=\"text-align:center\">✓</td> <td>Model name of the device (optional for web).</td> </tr> <tr> <td>Device-Platform</td> <td style=\"text-align:center\">✓</td> <td>The device's platform. Values are <code>android</code>, <code>ios</code>, or <code>web</code>.</td> </tr> <tr> <td>User-Agent</td> <td style=\"text-align:center\"></td> <td>The user agent of the client accessing the API.</td> </tr> </tbody> </table> <h4>HTTP Response Status Codes</h4> <table> <thead> <tr> <th style=\"text-align:center\"><strong>Code</strong></th> <th><strong>Description</strong></th> </tr> </thead> <tbody> <tr> <td style=\"text-align:center\">200</td> <td>OK, request proceed without error.</td> </tr> <tr> <td style=\"text-align:center\">204</td> <td>Request proceed successfully and is not returning any content (empty <code>data</code>).</td> </tr> <tr> <td style=\"text-align:center\">400</td> <td>Bad request, either the request header or the API parameter validation failed.</td> </tr> <tr> <td style=\"text-align:center\">401</td> <td>Unauthorized access, either the HTTP request header <code>Authorization</code> is missing or invalid.</td> </tr> <tr> <td style=\"text-align:center\">403</td> <td>Forbidden, the user does not have access to the requested resource.</td> </tr> <tr> <td style=\"text-align:center\">404</td> <td>Not found, requested resource does not exist.</td> </tr> <tr> <td style=\"text-align:center\">429</td> <td>Too many requests sent in a given amount of time, intended for use with rate-limiting schemes.</td> </tr> <tr> <td style=\"text-align:center\">491</td> <td>The API key specified in HTTP request header <code>API-Key</code> is invalid.</td> </tr> <tr> <td style=\"text-align:center\">500</td> <td>Internal server error while processing the request.</td> </tr> </tbody> </table> <h4>API Error Codes</h4> <table> <thead> <tr> <th style=\"text-align:center\"><strong>Code</strong></th> <th><strong>Description</strong></th> </tr> </thead> <tbody> <tr> <td style=\"text-align:center\">40001</td> <td>HTTP request header validation failed.</td> </tr> <tr> <td style=\"text-align:center\">40002</td> <td>API request parameter validation failed.</td> </tr> <tr> <td style=\"text-align:center\">40401</td> <td>The requested resource is not found.</td> </tr> <tr> <td style=\"text-align:center\">49101</td> <td>The API key is not provided.</td> </tr> <tr> <td style=\"text-align:center\">49102</td> <td>Failed to parse the API key, or the API key is invalid.</td> </tr> <tr> <td style=\"text-align:center\">49103</td> <td>The provided API key is not found.</td> </tr> <tr> <td style=\"text-align:center\">49104</td> <td>The provided API key is not intended to be used with the client's app platform.</td> </tr> <tr> <td style=\"text-align:center\">49105</td> <td>The provided API key is not intended to be used with the client's app identifier.</td> </tr> <tr> <td style=\"text-align:center\">49106</td> <td>The API key has expired.</td> </tr> <tr> <td style=\"text-align:center\">49107</td> <td>The API key is disabled.</td> </tr> <tr> <td style=\"text-align:center\">50001</td> <td>An error occurred while validating the API key.</td> </tr> <tr> <td style=\"text-align:center\">99999</td> <td>Other errors, usually without specific reason or action.</td> </tr> </tbody> </table>"
  },
  {
    "type": "post",
    "url": "/v1/client/account_verification/submit",
    "title": "Account Verification - Submit",
    "version": "1.0.0",
    "name": "AccountVerification_Submit",
    "group": "ClientAPI",
    "permission": [
      {
        "name": "client"
      }
    ],
    "description": "<p>Submit code for account verification.</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "long",
            "optional": false,
            "field": "otpID",
            "description": "<p>The OTP ID.</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "otpKey",
            "description": "<p>The OTP key.</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "otpCode",
            "description": "<p>The OTP code.</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "email",
            "description": "<p>The email address of the account to be verified.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Request Example:",
          "content": "{\n  \"otpID\": 128,\n  \"otpKey\": \"49ff390684515734c4c645df3884ed7c\",\n  \"otpCode\": \"123456\",\n  \"email\": \"john@doe.com\"\n}",
          "type": "json"
        }
      ]
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "boolean",
            "optional": false,
            "field": "success",
            "description": "<p>If the account verification is successful.</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "message",
            "description": "<p>The message.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success Response:",
          "content": "HTTP/1.1 200 OK\n{\n  \"status\": 200,\n  \"error\": {\n    \"code\": \"\",\n    \"message\": \"\",\n    \"field\": \"\"\n  },\n  \"data\": {\n    \"success\": true,\n    \"message\": \"Your account verification is successful\"\n  }\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "ParamValidationFailed",
            "description": "<p>The parameter validation failed.</p>"
          },
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "EmailNotRegistered",
            "description": "<p>The email address is not registered.</p>"
          },
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "HeaderValidationFailed",
            "description": "<p>The request header validation failed.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "ParamValidationFailed:",
          "content": "HTTP/1.1 200 OK\n{\n  \"status\": 400,\n  \"error\": {\n    \"code\": \"40002\",\n    \"message\": \"Email address format is invalid\",\n    \"field\": \"email\"\n  },\n  \"data\": {}\n}",
          "type": "json"
        },
        {
          "title": "EmailNotRegistered:",
          "content": "HTTP/1.1 200 OK\n{\n  \"status\": 400,\n  \"error\": {\n    \"code\": \"40002\",\n    \"message\": \"The email address is not registered\",\n    \"field\": \"email\"\n  },\n  \"data\": {}\n}",
          "type": "json"
        },
        {
          "title": "HeaderValidationFailed:",
          "content": "HTTP/1.1 200 OK\n{\n  \"status\": 400,\n  \"error\": {\n    \"code\": \"40001\",\n    \"message\": \"Request headers are required (API-Key, Device-Platform)\",\n    \"field\": \"\"\n  },\n  \"data\": {}\n}",
          "type": "json"
        }
      ]
    },
    "filename": "internal/app/basego-api/v1/endpoint/clientapi/account-verification-submit.go",
    "groupTitle": "Client API",
    "groupDescription": "<h4>HTTP Request Headers</h4> <table> <thead> <tr> <th><strong>Header Name</strong></th> <th style=\"text-align:center\"><strong>Required</strong></th> <th><strong>Description</strong></th> </tr> </thead> <tbody> <tr> <td>API-Key</td> <td style=\"text-align:center\">✓</td> <td>API key for accessing the API.</td> </tr> <tr> <td>App-Identifier</td> <td style=\"text-align:center\"></td> <td>The app's identifier (package name for Android, bundle ID for iOS, or origin URL for web).</td> </tr> <tr> <td>Content-Type</td> <td style=\"text-align:center\"></td> <td>Content type of the request body.</td> </tr> <tr> <td>Device-Identifier</td> <td style=\"text-align:center\">✓</td> <td>The device ID (optional for web).</td> </tr> <tr> <td>Device-Model</td> <td style=\"text-align:center\">✓</td> <td>Model name of the device (optional for web).</td> </tr> <tr> <td>Device-Platform</td> <td style=\"text-align:center\">✓</td> <td>The device's platform. Values are <code>android</code>, <code>ios</code>, or <code>web</code>.</td> </tr> <tr> <td>User-Agent</td> <td style=\"text-align:center\"></td> <td>The user agent of the client accessing the API.</td> </tr> </tbody> </table> <h4>HTTP Response Status Codes</h4> <table> <thead> <tr> <th style=\"text-align:center\"><strong>Code</strong></th> <th><strong>Description</strong></th> </tr> </thead> <tbody> <tr> <td style=\"text-align:center\">200</td> <td>OK, request proceed without error.</td> </tr> <tr> <td style=\"text-align:center\">204</td> <td>Request proceed successfully and is not returning any content (empty <code>data</code>).</td> </tr> <tr> <td style=\"text-align:center\">400</td> <td>Bad request, either the request header or the API parameter validation failed.</td> </tr> <tr> <td style=\"text-align:center\">401</td> <td>Unauthorized access, either the HTTP request header <code>Authorization</code> is missing or invalid.</td> </tr> <tr> <td style=\"text-align:center\">403</td> <td>Forbidden, the user does not have access to the requested resource.</td> </tr> <tr> <td style=\"text-align:center\">404</td> <td>Not found, requested resource does not exist.</td> </tr> <tr> <td style=\"text-align:center\">429</td> <td>Too many requests sent in a given amount of time, intended for use with rate-limiting schemes.</td> </tr> <tr> <td style=\"text-align:center\">491</td> <td>The API key specified in HTTP request header <code>API-Key</code> is invalid.</td> </tr> <tr> <td style=\"text-align:center\">500</td> <td>Internal server error while processing the request.</td> </tr> </tbody> </table> <h4>API Error Codes</h4> <table> <thead> <tr> <th style=\"text-align:center\"><strong>Code</strong></th> <th><strong>Description</strong></th> </tr> </thead> <tbody> <tr> <td style=\"text-align:center\">40001</td> <td>HTTP request header validation failed.</td> </tr> <tr> <td style=\"text-align:center\">40002</td> <td>API request parameter validation failed.</td> </tr> <tr> <td style=\"text-align:center\">40401</td> <td>The requested resource is not found.</td> </tr> <tr> <td style=\"text-align:center\">49101</td> <td>The API key is not provided.</td> </tr> <tr> <td style=\"text-align:center\">49102</td> <td>Failed to parse the API key, or the API key is invalid.</td> </tr> <tr> <td style=\"text-align:center\">49103</td> <td>The provided API key is not found.</td> </tr> <tr> <td style=\"text-align:center\">49104</td> <td>The provided API key is not intended to be used with the client's app platform.</td> </tr> <tr> <td style=\"text-align:center\">49105</td> <td>The provided API key is not intended to be used with the client's app identifier.</td> </tr> <tr> <td style=\"text-align:center\">49106</td> <td>The API key has expired.</td> </tr> <tr> <td style=\"text-align:center\">49107</td> <td>The API key is disabled.</td> </tr> <tr> <td style=\"text-align:center\">50001</td> <td>An error occurred while validating the API key.</td> </tr> <tr> <td style=\"text-align:center\">99999</td> <td>Other errors, usually without specific reason or action.</td> </tr> </tbody> </table>"
  },
  {
    "type": "post",
    "url": "/v1/client/register",
    "title": "Register",
    "version": "1.2.2",
    "name": "Register",
    "group": "ClientAPI",
    "permission": [
      {
        "name": "client"
      }
    ],
    "description": "<p>Register a new customer account.</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "fullName",
            "description": "<p>The full name.</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "email",
            "description": "<p>The email address.</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "password",
            "description": "<p>The password.</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": true,
            "field": "isTOSAccepted",
            "description": "<p>If the Terms of Service is accepted.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Request Example:",
          "content": "{\n  \"fullName\": \"John Doe\",\n  \"email\": \"john@doe.com\",\n  \"password\": \"this_is_password\",\n  \"isTOSAccepted\": true\n}",
          "type": "json"
        }
      ]
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "boolean",
            "optional": false,
            "field": "success",
            "description": "<p>If the registration is successful.</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "message",
            "description": "<p>The message, if failed.</p>"
          },
          {
            "group": "Success 200",
            "type": "long",
            "optional": false,
            "field": "otpID",
            "description": "<p>The OTP ID for OTP verification.</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "otpKey",
            "description": "<p>The OTP key.</p>"
          },
          {
            "group": "Success 200",
            "type": "integer",
            "optional": false,
            "field": "codeLength",
            "description": "<p>The OTP code's length.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success Response:",
          "content": "HTTP/1.1 200 OK\n{\n  \"status\": 200,\n  \"error\": {\n    \"code\": \"\",\n    \"message\": \"\",\n    \"field\": \"\"\n  },\n  \"data\": {\n    \"success\": true,\n    \"message\": \"\",\n    \"otpID\": 128,\n    \"otpKey\": \"49ff390684515734c4c645df3884ed7c\",\n    \"codeLength\": 6\n  }\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "ParamValidationFailed",
            "description": "<p>The parameter validation failed.</p>"
          },
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "EmailAlreadyRegistered",
            "description": "<p>The email address is already registered.</p>"
          },
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "HeaderValidationFailed",
            "description": "<p>The request header validation failed.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "ParamValidationFailed:",
          "content": "HTTP/1.1 200 OK\n{\n  \"status\": 400,\n  \"error\": {\n    \"code\": \"40002\",\n    \"message\": \"Email address format is invalid\",\n    \"field\": \"email\"\n  },\n  \"data\": {}\n}",
          "type": "json"
        },
        {
          "title": "EmailAlreadyRegistered:",
          "content": "HTTP/1.1 200 OK\n{\n  \"status\": 400,\n  \"error\": {\n    \"code\": \"40002\",\n    \"message\": \"The email address is already registered\",\n    \"field\": \"email\"\n  },\n  \"data\": {}\n}",
          "type": "json"
        },
        {
          "title": "HeaderValidationFailed:",
          "content": "HTTP/1.1 200 OK\n{\n  \"status\": 400,\n  \"error\": {\n    \"code\": \"40001\",\n    \"message\": \"Request headers are required (API-Key, Device-Platform)\",\n    \"field\": \"\"\n  },\n  \"data\": {}\n}",
          "type": "json"
        }
      ]
    },
    "filename": "internal/app/basego-api/v1/endpoint/clientapi/register.go",
    "groupTitle": "Client API",
    "groupDescription": "<h4>HTTP Request Headers</h4> <table> <thead> <tr> <th><strong>Header Name</strong></th> <th style=\"text-align:center\"><strong>Required</strong></th> <th><strong>Description</strong></th> </tr> </thead> <tbody> <tr> <td>API-Key</td> <td style=\"text-align:center\">✓</td> <td>API key for accessing the API.</td> </tr> <tr> <td>App-Identifier</td> <td style=\"text-align:center\"></td> <td>The app's identifier (package name for Android, bundle ID for iOS, or origin URL for web).</td> </tr> <tr> <td>Content-Type</td> <td style=\"text-align:center\"></td> <td>Content type of the request body.</td> </tr> <tr> <td>Device-Identifier</td> <td style=\"text-align:center\">✓</td> <td>The device ID (optional for web).</td> </tr> <tr> <td>Device-Model</td> <td style=\"text-align:center\">✓</td> <td>Model name of the device (optional for web).</td> </tr> <tr> <td>Device-Platform</td> <td style=\"text-align:center\">✓</td> <td>The device's platform. Values are <code>android</code>, <code>ios</code>, or <code>web</code>.</td> </tr> <tr> <td>User-Agent</td> <td style=\"text-align:center\"></td> <td>The user agent of the client accessing the API.</td> </tr> </tbody> </table> <h4>HTTP Response Status Codes</h4> <table> <thead> <tr> <th style=\"text-align:center\"><strong>Code</strong></th> <th><strong>Description</strong></th> </tr> </thead> <tbody> <tr> <td style=\"text-align:center\">200</td> <td>OK, request proceed without error.</td> </tr> <tr> <td style=\"text-align:center\">204</td> <td>Request proceed successfully and is not returning any content (empty <code>data</code>).</td> </tr> <tr> <td style=\"text-align:center\">400</td> <td>Bad request, either the request header or the API parameter validation failed.</td> </tr> <tr> <td style=\"text-align:center\">401</td> <td>Unauthorized access, either the HTTP request header <code>Authorization</code> is missing or invalid.</td> </tr> <tr> <td style=\"text-align:center\">403</td> <td>Forbidden, the user does not have access to the requested resource.</td> </tr> <tr> <td style=\"text-align:center\">404</td> <td>Not found, requested resource does not exist.</td> </tr> <tr> <td style=\"text-align:center\">429</td> <td>Too many requests sent in a given amount of time, intended for use with rate-limiting schemes.</td> </tr> <tr> <td style=\"text-align:center\">491</td> <td>The API key specified in HTTP request header <code>API-Key</code> is invalid.</td> </tr> <tr> <td style=\"text-align:center\">500</td> <td>Internal server error while processing the request.</td> </tr> </tbody> </table> <h4>API Error Codes</h4> <table> <thead> <tr> <th style=\"text-align:center\"><strong>Code</strong></th> <th><strong>Description</strong></th> </tr> </thead> <tbody> <tr> <td style=\"text-align:center\">40001</td> <td>HTTP request header validation failed.</td> </tr> <tr> <td style=\"text-align:center\">40002</td> <td>API request parameter validation failed.</td> </tr> <tr> <td style=\"text-align:center\">40401</td> <td>The requested resource is not found.</td> </tr> <tr> <td style=\"text-align:center\">49101</td> <td>The API key is not provided.</td> </tr> <tr> <td style=\"text-align:center\">49102</td> <td>Failed to parse the API key, or the API key is invalid.</td> </tr> <tr> <td style=\"text-align:center\">49103</td> <td>The provided API key is not found.</td> </tr> <tr> <td style=\"text-align:center\">49104</td> <td>The provided API key is not intended to be used with the client's app platform.</td> </tr> <tr> <td style=\"text-align:center\">49105</td> <td>The provided API key is not intended to be used with the client's app identifier.</td> </tr> <tr> <td style=\"text-align:center\">49106</td> <td>The API key has expired.</td> </tr> <tr> <td style=\"text-align:center\">49107</td> <td>The API key is disabled.</td> </tr> <tr> <td style=\"text-align:center\">50001</td> <td>An error occurred while validating the API key.</td> </tr> <tr> <td style=\"text-align:center\">99999</td> <td>Other errors, usually without specific reason or action.</td> </tr> </tbody> </table>"
  },
  {
    "type": "post",
    "url": "/v1/client/register",
    "title": "Register",
    "version": "1.0.0",
    "name": "Register",
    "group": "ClientAPI",
    "permission": [
      {
        "name": "client"
      }
    ],
    "description": "<p>Register a new customer account.</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "fullName",
            "description": "<p>The full name.</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "email",
            "description": "<p>The email address.</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "password",
            "description": "<p>The password.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Request Example:",
          "content": "{\n  \"fullName\": \"John Doe\",\n  \"email\": \"john@doe.com\",\n  \"password\": \"this_is_password\"\n}",
          "type": "json"
        }
      ]
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "boolean",
            "optional": false,
            "field": "success",
            "description": "<p>If the registration is successful.</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "message",
            "description": "<p>The message, if failed.</p>"
          },
          {
            "group": "Success 200",
            "type": "long",
            "optional": false,
            "field": "otpID",
            "description": "<p>The OTP ID for OTP verification.</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "otpKey",
            "description": "<p>The OTP key.</p>"
          },
          {
            "group": "Success 200",
            "type": "integer",
            "optional": false,
            "field": "codeLength",
            "description": "<p>The OTP code's length.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success Response:",
          "content": "HTTP/1.1 200 OK\n{\n  \"status\": 200,\n  \"error\": {\n    \"code\": \"\",\n    \"message\": \"\",\n    \"field\": \"\"\n  },\n  \"data\": {\n    \"success\": true,\n    \"message\": \"\",\n    \"otpID\": 128,\n    \"otpKey\": \"49ff390684515734c4c645df3884ed7c\",\n    \"codeLength\": 6\n  }\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "ParamValidationFailed",
            "description": "<p>The parameter validation failed.</p>"
          },
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "EmailAlreadyRegistered",
            "description": "<p>The email address is already registered.</p>"
          },
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "HeaderValidationFailed",
            "description": "<p>The request header validation failed.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "ParamValidationFailed:",
          "content": "HTTP/1.1 200 OK\n{\n  \"status\": 400,\n  \"error\": {\n    \"code\": \"40002\",\n    \"message\": \"Email address format is invalid\",\n    \"field\": \"email\"\n  },\n  \"data\": {}\n}",
          "type": "json"
        },
        {
          "title": "EmailAlreadyRegistered:",
          "content": "HTTP/1.1 200 OK\n{\n  \"status\": 400,\n  \"error\": {\n    \"code\": \"40002\",\n    \"message\": \"The email address is already registered\",\n    \"field\": \"email\"\n  },\n  \"data\": {}\n}",
          "type": "json"
        },
        {
          "title": "HeaderValidationFailed:",
          "content": "HTTP/1.1 200 OK\n{\n  \"status\": 400,\n  \"error\": {\n    \"code\": \"40001\",\n    \"message\": \"Request headers are required (API-Key, Device-Platform)\",\n    \"field\": \"\"\n  },\n  \"data\": {}\n}",
          "type": "json"
        }
      ]
    },
    "filename": "internal/app/basego-api/v1/endpoint/clientapi/register.go",
    "groupTitle": "Client API",
    "groupDescription": "<h4>HTTP Request Headers</h4> <table> <thead> <tr> <th><strong>Header Name</strong></th> <th style=\"text-align:center\"><strong>Required</strong></th> <th><strong>Description</strong></th> </tr> </thead> <tbody> <tr> <td>API-Key</td> <td style=\"text-align:center\">✓</td> <td>API key for accessing the API.</td> </tr> <tr> <td>App-Identifier</td> <td style=\"text-align:center\"></td> <td>The app's identifier (package name for Android, bundle ID for iOS, or origin URL for web).</td> </tr> <tr> <td>Content-Type</td> <td style=\"text-align:center\"></td> <td>Content type of the request body.</td> </tr> <tr> <td>Device-Identifier</td> <td style=\"text-align:center\">✓</td> <td>The device ID (optional for web).</td> </tr> <tr> <td>Device-Model</td> <td style=\"text-align:center\">✓</td> <td>Model name of the device (optional for web).</td> </tr> <tr> <td>Device-Platform</td> <td style=\"text-align:center\">✓</td> <td>The device's platform. Values are <code>android</code>, <code>ios</code>, or <code>web</code>.</td> </tr> <tr> <td>User-Agent</td> <td style=\"text-align:center\"></td> <td>The user agent of the client accessing the API.</td> </tr> </tbody> </table> <h4>HTTP Response Status Codes</h4> <table> <thead> <tr> <th style=\"text-align:center\"><strong>Code</strong></th> <th><strong>Description</strong></th> </tr> </thead> <tbody> <tr> <td style=\"text-align:center\">200</td> <td>OK, request proceed without error.</td> </tr> <tr> <td style=\"text-align:center\">204</td> <td>Request proceed successfully and is not returning any content (empty <code>data</code>).</td> </tr> <tr> <td style=\"text-align:center\">400</td> <td>Bad request, either the request header or the API parameter validation failed.</td> </tr> <tr> <td style=\"text-align:center\">401</td> <td>Unauthorized access, either the HTTP request header <code>Authorization</code> is missing or invalid.</td> </tr> <tr> <td style=\"text-align:center\">403</td> <td>Forbidden, the user does not have access to the requested resource.</td> </tr> <tr> <td style=\"text-align:center\">404</td> <td>Not found, requested resource does not exist.</td> </tr> <tr> <td style=\"text-align:center\">429</td> <td>Too many requests sent in a given amount of time, intended for use with rate-limiting schemes.</td> </tr> <tr> <td style=\"text-align:center\">491</td> <td>The API key specified in HTTP request header <code>API-Key</code> is invalid.</td> </tr> <tr> <td style=\"text-align:center\">500</td> <td>Internal server error while processing the request.</td> </tr> </tbody> </table> <h4>API Error Codes</h4> <table> <thead> <tr> <th style=\"text-align:center\"><strong>Code</strong></th> <th><strong>Description</strong></th> </tr> </thead> <tbody> <tr> <td style=\"text-align:center\">40001</td> <td>HTTP request header validation failed.</td> </tr> <tr> <td style=\"text-align:center\">40002</td> <td>API request parameter validation failed.</td> </tr> <tr> <td style=\"text-align:center\">40401</td> <td>The requested resource is not found.</td> </tr> <tr> <td style=\"text-align:center\">49101</td> <td>The API key is not provided.</td> </tr> <tr> <td style=\"text-align:center\">49102</td> <td>Failed to parse the API key, or the API key is invalid.</td> </tr> <tr> <td style=\"text-align:center\">49103</td> <td>The provided API key is not found.</td> </tr> <tr> <td style=\"text-align:center\">49104</td> <td>The provided API key is not intended to be used with the client's app platform.</td> </tr> <tr> <td style=\"text-align:center\">49105</td> <td>The provided API key is not intended to be used with the client's app identifier.</td> </tr> <tr> <td style=\"text-align:center\">49106</td> <td>The API key has expired.</td> </tr> <tr> <td style=\"text-align:center\">49107</td> <td>The API key is disabled.</td> </tr> <tr> <td style=\"text-align:center\">50001</td> <td>An error occurred while validating the API key.</td> </tr> <tr> <td style=\"text-align:center\">99999</td> <td>Other errors, usually without specific reason or action.</td> </tr> </tbody> </table>"
  },
  {
    "type": "post",
    "url": "/v1/client/reset_password/request_token",
    "title": "Reset Password - Request Token",
    "version": "1.0.0",
    "name": "ResetPassword_RequestToken",
    "group": "ClientAPI",
    "permission": [
      {
        "name": "client"
      }
    ],
    "description": "<p>Request token for reset password.</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "email",
            "description": "<p>The account's email address.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Request Example:",
          "content": "{\n  \"email\": \"john@doe.com\",\n}",
          "type": "json"
        }
      ]
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "boolean",
            "optional": false,
            "field": "success",
            "description": "<p>If email containing the reset password token is sent successfully.</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "message",
            "description": "<p>The message.</p>"
          },
          {
            "group": "Success 200",
            "type": "long",
            "optional": false,
            "field": "otpID",
            "description": "<p>The OTP ID for reset paassword token.</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "otpKey",
            "description": "<p>The OTP key.</p>"
          },
          {
            "group": "Success 200",
            "type": "integer",
            "optional": false,
            "field": "codeLength",
            "description": "<p>The OTP code's length.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success Response:",
          "content": "HTTP/1.1 200 OK\n{\n  \"status\": 200,\n  \"error\": {\n    \"code\": \"\",\n    \"message\": \"\",\n    \"field\": \"\"\n  },\n  \"data\": {\n    \"success\": true,\n    \"message\": \"\",\n    \"otpID\": 128,\n    \"otpKey\": \"49ff390684515734c4c645df3884ed7c\",\n    \"codeLength\": 6\n  }\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "ParamValidationFailed",
            "description": "<p>The parameter validation failed.</p>"
          },
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "HeaderValidationFailed",
            "description": "<p>The request header validation failed.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "ParamValidationFailed:",
          "content": "HTTP/1.1 200 OK\n{\n  \"status\": 400,\n  \"error\": {\n    \"code\": \"40002\",\n    \"message\": \"Email address is invalid\",\n    \"field\": \"email\"\n  },\n  \"data\": {}\n}",
          "type": "json"
        },
        {
          "title": "HeaderValidationFailed:",
          "content": "HTTP/1.1 200 OK\n{\n  \"status\": 400,\n  \"error\": {\n    \"code\": \"40001\",\n    \"message\": \"Request headers are required (API-Key, Device-Platform)\",\n    \"field\": \"\"\n  },\n  \"data\": {}\n}",
          "type": "json"
        }
      ]
    },
    "filename": "internal/app/basego-api/v1/endpoint/clientapi/reset-password-request-token.go",
    "groupTitle": "Client API",
    "groupDescription": "<h4>HTTP Request Headers</h4> <table> <thead> <tr> <th><strong>Header Name</strong></th> <th style=\"text-align:center\"><strong>Required</strong></th> <th><strong>Description</strong></th> </tr> </thead> <tbody> <tr> <td>API-Key</td> <td style=\"text-align:center\">✓</td> <td>API key for accessing the API.</td> </tr> <tr> <td>App-Identifier</td> <td style=\"text-align:center\"></td> <td>The app's identifier (package name for Android, bundle ID for iOS, or origin URL for web).</td> </tr> <tr> <td>Content-Type</td> <td style=\"text-align:center\"></td> <td>Content type of the request body.</td> </tr> <tr> <td>Device-Identifier</td> <td style=\"text-align:center\">✓</td> <td>The device ID (optional for web).</td> </tr> <tr> <td>Device-Model</td> <td style=\"text-align:center\">✓</td> <td>Model name of the device (optional for web).</td> </tr> <tr> <td>Device-Platform</td> <td style=\"text-align:center\">✓</td> <td>The device's platform. Values are <code>android</code>, <code>ios</code>, or <code>web</code>.</td> </tr> <tr> <td>User-Agent</td> <td style=\"text-align:center\"></td> <td>The user agent of the client accessing the API.</td> </tr> </tbody> </table> <h4>HTTP Response Status Codes</h4> <table> <thead> <tr> <th style=\"text-align:center\"><strong>Code</strong></th> <th><strong>Description</strong></th> </tr> </thead> <tbody> <tr> <td style=\"text-align:center\">200</td> <td>OK, request proceed without error.</td> </tr> <tr> <td style=\"text-align:center\">204</td> <td>Request proceed successfully and is not returning any content (empty <code>data</code>).</td> </tr> <tr> <td style=\"text-align:center\">400</td> <td>Bad request, either the request header or the API parameter validation failed.</td> </tr> <tr> <td style=\"text-align:center\">401</td> <td>Unauthorized access, either the HTTP request header <code>Authorization</code> is missing or invalid.</td> </tr> <tr> <td style=\"text-align:center\">403</td> <td>Forbidden, the user does not have access to the requested resource.</td> </tr> <tr> <td style=\"text-align:center\">404</td> <td>Not found, requested resource does not exist.</td> </tr> <tr> <td style=\"text-align:center\">429</td> <td>Too many requests sent in a given amount of time, intended for use with rate-limiting schemes.</td> </tr> <tr> <td style=\"text-align:center\">491</td> <td>The API key specified in HTTP request header <code>API-Key</code> is invalid.</td> </tr> <tr> <td style=\"text-align:center\">500</td> <td>Internal server error while processing the request.</td> </tr> </tbody> </table> <h4>API Error Codes</h4> <table> <thead> <tr> <th style=\"text-align:center\"><strong>Code</strong></th> <th><strong>Description</strong></th> </tr> </thead> <tbody> <tr> <td style=\"text-align:center\">40001</td> <td>HTTP request header validation failed.</td> </tr> <tr> <td style=\"text-align:center\">40002</td> <td>API request parameter validation failed.</td> </tr> <tr> <td style=\"text-align:center\">40401</td> <td>The requested resource is not found.</td> </tr> <tr> <td style=\"text-align:center\">49101</td> <td>The API key is not provided.</td> </tr> <tr> <td style=\"text-align:center\">49102</td> <td>Failed to parse the API key, or the API key is invalid.</td> </tr> <tr> <td style=\"text-align:center\">49103</td> <td>The provided API key is not found.</td> </tr> <tr> <td style=\"text-align:center\">49104</td> <td>The provided API key is not intended to be used with the client's app platform.</td> </tr> <tr> <td style=\"text-align:center\">49105</td> <td>The provided API key is not intended to be used with the client's app identifier.</td> </tr> <tr> <td style=\"text-align:center\">49106</td> <td>The API key has expired.</td> </tr> <tr> <td style=\"text-align:center\">49107</td> <td>The API key is disabled.</td> </tr> <tr> <td style=\"text-align:center\">50001</td> <td>An error occurred while validating the API key.</td> </tr> <tr> <td style=\"text-align:center\">99999</td> <td>Other errors, usually without specific reason or action.</td> </tr> </tbody> </table>"
  },
  {
    "type": "post",
    "url": "/v1/client/reset_password/set_password",
    "title": "Reset Password - Set Password",
    "version": "1.0.0",
    "name": "ResetPassword_SetPassword",
    "group": "ClientAPI",
    "permission": [
      {
        "name": "client"
      }
    ],
    "description": "<p>Set new password for a customer account.</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "long",
            "optional": false,
            "field": "otpID",
            "description": "<p>The OTP ID of the reset password token.</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "otpKey",
            "description": "<p>The OTP key.</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "otpCode",
            "description": "<p>The OTP code.</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "email",
            "description": "<p>The account's email address.</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "password",
            "description": "<p>The new password.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Request Example:",
          "content": "{\n  \"otpID\": 128,\n  \"otpKey\": \"49ff390684515734c4c645df3884ed7c\",\n  \"otpCode\": \"123456\",\n  \"email\": \"john@doe.com\",\n  \"password\": \"This_is_New_Password_123\"\n}",
          "type": "json"
        }
      ]
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "boolean",
            "optional": false,
            "field": "success",
            "description": "<p>If the password is changed successfully.</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "message",
            "description": "<p>The message.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success Response:",
          "content": "HTTP/1.1 200 OK\n{\n  \"status\": 200,\n  \"error\": {\n    \"code\": \"\",\n    \"message\": \"\",\n    \"field\": \"\"\n  },\n  \"data\": {\n    \"success\": true,\n    \"message\": \"Password changed successfully\",\n  }\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "ParamValidationFailed",
            "description": "<p>The parameter validation failed.</p>"
          },
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "PasswordFormatInvalid",
            "description": "<p>The new password's format is invalid.</p>"
          },
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "HeaderValidationFailed",
            "description": "<p>The request header validation failed.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "ParamValidationFailed:",
          "content": "HTTP/1.1 200 OK\n{\n  \"status\": 400,\n  \"error\": {\n    \"code\": \"40002\",\n    \"message\": \"Password is required\",\n    \"field\": \"password\"\n  },\n  \"data\": {}\n}",
          "type": "json"
        },
        {
          "title": "PasswordFormatInvalid:",
          "content": "HTTP/1.1 200 OK\n{\n  \"status\": 400,\n  \"error\": {\n    \"code\": \"40002\",\n    \"message\": \"Password must contain at least 1 lowercase, uppercase, and special characters and 1 number\",\n    \"field\": \"password\"\n  },\n  \"data\": {}\n}",
          "type": "json"
        },
        {
          "title": "HeaderValidationFailed:",
          "content": "HTTP/1.1 200 OK\n{\n  \"status\": 400,\n  \"error\": {\n    \"code\": \"40001\",\n    \"message\": \"Request headers are required (API-Key, Device-Platform)\",\n    \"field\": \"\"\n  },\n  \"data\": {}\n}",
          "type": "json"
        }
      ]
    },
    "filename": "internal/app/basego-api/v1/endpoint/clientapi/reset-password-set-password.go",
    "groupTitle": "Client API",
    "groupDescription": "<h4>HTTP Request Headers</h4> <table> <thead> <tr> <th><strong>Header Name</strong></th> <th style=\"text-align:center\"><strong>Required</strong></th> <th><strong>Description</strong></th> </tr> </thead> <tbody> <tr> <td>API-Key</td> <td style=\"text-align:center\">✓</td> <td>API key for accessing the API.</td> </tr> <tr> <td>App-Identifier</td> <td style=\"text-align:center\"></td> <td>The app's identifier (package name for Android, bundle ID for iOS, or origin URL for web).</td> </tr> <tr> <td>Content-Type</td> <td style=\"text-align:center\"></td> <td>Content type of the request body.</td> </tr> <tr> <td>Device-Identifier</td> <td style=\"text-align:center\">✓</td> <td>The device ID (optional for web).</td> </tr> <tr> <td>Device-Model</td> <td style=\"text-align:center\">✓</td> <td>Model name of the device (optional for web).</td> </tr> <tr> <td>Device-Platform</td> <td style=\"text-align:center\">✓</td> <td>The device's platform. Values are <code>android</code>, <code>ios</code>, or <code>web</code>.</td> </tr> <tr> <td>User-Agent</td> <td style=\"text-align:center\"></td> <td>The user agent of the client accessing the API.</td> </tr> </tbody> </table> <h4>HTTP Response Status Codes</h4> <table> <thead> <tr> <th style=\"text-align:center\"><strong>Code</strong></th> <th><strong>Description</strong></th> </tr> </thead> <tbody> <tr> <td style=\"text-align:center\">200</td> <td>OK, request proceed without error.</td> </tr> <tr> <td style=\"text-align:center\">204</td> <td>Request proceed successfully and is not returning any content (empty <code>data</code>).</td> </tr> <tr> <td style=\"text-align:center\">400</td> <td>Bad request, either the request header or the API parameter validation failed.</td> </tr> <tr> <td style=\"text-align:center\">401</td> <td>Unauthorized access, either the HTTP request header <code>Authorization</code> is missing or invalid.</td> </tr> <tr> <td style=\"text-align:center\">403</td> <td>Forbidden, the user does not have access to the requested resource.</td> </tr> <tr> <td style=\"text-align:center\">404</td> <td>Not found, requested resource does not exist.</td> </tr> <tr> <td style=\"text-align:center\">429</td> <td>Too many requests sent in a given amount of time, intended for use with rate-limiting schemes.</td> </tr> <tr> <td style=\"text-align:center\">491</td> <td>The API key specified in HTTP request header <code>API-Key</code> is invalid.</td> </tr> <tr> <td style=\"text-align:center\">500</td> <td>Internal server error while processing the request.</td> </tr> </tbody> </table> <h4>API Error Codes</h4> <table> <thead> <tr> <th style=\"text-align:center\"><strong>Code</strong></th> <th><strong>Description</strong></th> </tr> </thead> <tbody> <tr> <td style=\"text-align:center\">40001</td> <td>HTTP request header validation failed.</td> </tr> <tr> <td style=\"text-align:center\">40002</td> <td>API request parameter validation failed.</td> </tr> <tr> <td style=\"text-align:center\">40401</td> <td>The requested resource is not found.</td> </tr> <tr> <td style=\"text-align:center\">49101</td> <td>The API key is not provided.</td> </tr> <tr> <td style=\"text-align:center\">49102</td> <td>Failed to parse the API key, or the API key is invalid.</td> </tr> <tr> <td style=\"text-align:center\">49103</td> <td>The provided API key is not found.</td> </tr> <tr> <td style=\"text-align:center\">49104</td> <td>The provided API key is not intended to be used with the client's app platform.</td> </tr> <tr> <td style=\"text-align:center\">49105</td> <td>The provided API key is not intended to be used with the client's app identifier.</td> </tr> <tr> <td style=\"text-align:center\">49106</td> <td>The API key has expired.</td> </tr> <tr> <td style=\"text-align:center\">49107</td> <td>The API key is disabled.</td> </tr> <tr> <td style=\"text-align:center\">50001</td> <td>An error occurred while validating the API key.</td> </tr> <tr> <td style=\"text-align:center\">99999</td> <td>Other errors, usually without specific reason or action.</td> </tr> </tbody> </table>"
  },
  {
    "type": "post",
    "url": "/v1/client/reset_password/verify_token",
    "title": "Reset Password - Verify Token",
    "version": "1.0.0",
    "name": "ResetPassword_VerifyToken",
    "group": "ClientAPI",
    "permission": [
      {
        "name": "client"
      }
    ],
    "description": "<p>Verify a reset password token.</p>",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "long",
            "optional": false,
            "field": "otpID",
            "description": "<p>The OTP ID of the reset password token.</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "otpKey",
            "description": "<p>The OTP key.</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "otpCode",
            "description": "<p>The OTP code.</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "email",
            "description": "<p>The account's email address.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Request Example:",
          "content": "{\n  \"otpID\": 128,\n  \"otpKey\": \"49ff390684515734c4c645df3884ed7c\",\n  \"otpCode\": \"123456\",\n  \"email\": \"john@doe.com\"\n}",
          "type": "json"
        }
      ]
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "boolean",
            "optional": false,
            "field": "isValid",
            "description": "<p>If the reset password token is valid.</p>"
          },
          {
            "group": "Success 200",
            "type": "string",
            "optional": false,
            "field": "message",
            "description": "<p>The message.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success Response:",
          "content": "HTTP/1.1 200 OK\n{\n  \"status\": 200,\n  \"error\": {\n    \"code\": \"\",\n    \"message\": \"\",\n    \"field\": \"\"\n  },\n  \"data\": {\n    \"isValid\": true,\n    \"message\": \"\"\n  }\n}",
          "type": "json"
        }
      ]
    },
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "ParamValidationFailed",
            "description": "<p>The parameter validation failed.</p>"
          },
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "HeaderValidationFailed",
            "description": "<p>The request header validation failed.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "ParamValidationFailed:",
          "content": "HTTP/1.1 200 OK\n{\n  \"status\": 400,\n  \"error\": {\n    \"code\": \"40002\",\n    \"message\": \"Token is required\",\n    \"field\": \"otpCode\"\n  },\n  \"data\": {}\n}",
          "type": "json"
        },
        {
          "title": "HeaderValidationFailed:",
          "content": "HTTP/1.1 200 OK\n{\n  \"status\": 400,\n  \"error\": {\n    \"code\": \"40001\",\n    \"message\": \"Request headers are required (API-Key, Device-Platform)\",\n    \"field\": \"\"\n  },\n  \"data\": {}\n}",
          "type": "json"
        }
      ]
    },
    "filename": "internal/app/basego-api/v1/endpoint/clientapi/reset-password-verify-token.go",
    "groupTitle": "Client API",
    "groupDescription": "<h4>HTTP Request Headers</h4> <table> <thead> <tr> <th><strong>Header Name</strong></th> <th style=\"text-align:center\"><strong>Required</strong></th> <th><strong>Description</strong></th> </tr> </thead> <tbody> <tr> <td>API-Key</td> <td style=\"text-align:center\">✓</td> <td>API key for accessing the API.</td> </tr> <tr> <td>App-Identifier</td> <td style=\"text-align:center\"></td> <td>The app's identifier (package name for Android, bundle ID for iOS, or origin URL for web).</td> </tr> <tr> <td>Content-Type</td> <td style=\"text-align:center\"></td> <td>Content type of the request body.</td> </tr> <tr> <td>Device-Identifier</td> <td style=\"text-align:center\">✓</td> <td>The device ID (optional for web).</td> </tr> <tr> <td>Device-Model</td> <td style=\"text-align:center\">✓</td> <td>Model name of the device (optional for web).</td> </tr> <tr> <td>Device-Platform</td> <td style=\"text-align:center\">✓</td> <td>The device's platform. Values are <code>android</code>, <code>ios</code>, or <code>web</code>.</td> </tr> <tr> <td>User-Agent</td> <td style=\"text-align:center\"></td> <td>The user agent of the client accessing the API.</td> </tr> </tbody> </table> <h4>HTTP Response Status Codes</h4> <table> <thead> <tr> <th style=\"text-align:center\"><strong>Code</strong></th> <th><strong>Description</strong></th> </tr> </thead> <tbody> <tr> <td style=\"text-align:center\">200</td> <td>OK, request proceed without error.</td> </tr> <tr> <td style=\"text-align:center\">204</td> <td>Request proceed successfully and is not returning any content (empty <code>data</code>).</td> </tr> <tr> <td style=\"text-align:center\">400</td> <td>Bad request, either the request header or the API parameter validation failed.</td> </tr> <tr> <td style=\"text-align:center\">401</td> <td>Unauthorized access, either the HTTP request header <code>Authorization</code> is missing or invalid.</td> </tr> <tr> <td style=\"text-align:center\">403</td> <td>Forbidden, the user does not have access to the requested resource.</td> </tr> <tr> <td style=\"text-align:center\">404</td> <td>Not found, requested resource does not exist.</td> </tr> <tr> <td style=\"text-align:center\">429</td> <td>Too many requests sent in a given amount of time, intended for use with rate-limiting schemes.</td> </tr> <tr> <td style=\"text-align:center\">491</td> <td>The API key specified in HTTP request header <code>API-Key</code> is invalid.</td> </tr> <tr> <td style=\"text-align:center\">500</td> <td>Internal server error while processing the request.</td> </tr> </tbody> </table> <h4>API Error Codes</h4> <table> <thead> <tr> <th style=\"text-align:center\"><strong>Code</strong></th> <th><strong>Description</strong></th> </tr> </thead> <tbody> <tr> <td style=\"text-align:center\">40001</td> <td>HTTP request header validation failed.</td> </tr> <tr> <td style=\"text-align:center\">40002</td> <td>API request parameter validation failed.</td> </tr> <tr> <td style=\"text-align:center\">40401</td> <td>The requested resource is not found.</td> </tr> <tr> <td style=\"text-align:center\">49101</td> <td>The API key is not provided.</td> </tr> <tr> <td style=\"text-align:center\">49102</td> <td>Failed to parse the API key, or the API key is invalid.</td> </tr> <tr> <td style=\"text-align:center\">49103</td> <td>The provided API key is not found.</td> </tr> <tr> <td style=\"text-align:center\">49104</td> <td>The provided API key is not intended to be used with the client's app platform.</td> </tr> <tr> <td style=\"text-align:center\">49105</td> <td>The provided API key is not intended to be used with the client's app identifier.</td> </tr> <tr> <td style=\"text-align:center\">49106</td> <td>The API key has expired.</td> </tr> <tr> <td style=\"text-align:center\">49107</td> <td>The API key is disabled.</td> </tr> <tr> <td style=\"text-align:center\">50001</td> <td>An error occurred while validating the API key.</td> </tr> <tr> <td style=\"text-align:center\">99999</td> <td>Other errors, usually without specific reason or action.</td> </tr> </tbody> </table>"
  },
  {
    "type": "post",
    "url": "/v1/client/server_time",
    "title": "Get Server Time",
    "version": "1.0.0",
    "name": "ServerTime",
    "group": "ClientAPI",
    "permission": [
      {
        "name": "client"
      }
    ],
    "description": "<p>Get the server's current timestamp.</p>",
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "object",
            "optional": false,
            "field": "timestamp",
            "description": "<p>The server's current timestamp.</p>"
          },
          {
            "group": "Success 200",
            "type": "long",
            "optional": false,
            "field": "timestamp.seconds",
            "description": "<p>The timestamp in seconds.</p>"
          },
          {
            "group": "Success 200",
            "type": "long",
            "optional": false,
            "field": "timestamp.milliseconds",
            "description": "<p>The timestamp in milliseconds.</p>"
          },
          {
            "group": "Success 200",
            "type": "long",
            "optional": false,
            "field": "timestamp.nanoseconds",
            "description": "<p>The timestamp in nanoseconds.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success Response:",
          "content": "HTTP/1.1 200 OK\n{\n  \"status\": 200,\n  \"error\": {\n    \"code\": \"\",\n    \"message\": \"\",\n    \"field\": \"\"\n  },\n  \"data\": {\n    \"timestamp\": {\n      \"seconds\": 1551841641,\n      \"milliseconds\": 1551841641095,\n      \"nanoseconds\": 1551841641095244400\n    }\n  }\n}",
          "type": "json"
        }
      ]
    },
    "filename": "internal/app/basego-api/v1/endpoint/clientapi/server-time.go",
    "groupTitle": "Client API",
    "groupDescription": "<h4>HTTP Request Headers</h4> <table> <thead> <tr> <th><strong>Header Name</strong></th> <th style=\"text-align:center\"><strong>Required</strong></th> <th><strong>Description</strong></th> </tr> </thead> <tbody> <tr> <td>API-Key</td> <td style=\"text-align:center\">✓</td> <td>API key for accessing the API.</td> </tr> <tr> <td>App-Identifier</td> <td style=\"text-align:center\"></td> <td>The app's identifier (package name for Android, bundle ID for iOS, or origin URL for web).</td> </tr> <tr> <td>Content-Type</td> <td style=\"text-align:center\"></td> <td>Content type of the request body.</td> </tr> <tr> <td>Device-Identifier</td> <td style=\"text-align:center\">✓</td> <td>The device ID (optional for web).</td> </tr> <tr> <td>Device-Model</td> <td style=\"text-align:center\">✓</td> <td>Model name of the device (optional for web).</td> </tr> <tr> <td>Device-Platform</td> <td style=\"text-align:center\">✓</td> <td>The device's platform. Values are <code>android</code>, <code>ios</code>, or <code>web</code>.</td> </tr> <tr> <td>User-Agent</td> <td style=\"text-align:center\"></td> <td>The user agent of the client accessing the API.</td> </tr> </tbody> </table> <h4>HTTP Response Status Codes</h4> <table> <thead> <tr> <th style=\"text-align:center\"><strong>Code</strong></th> <th><strong>Description</strong></th> </tr> </thead> <tbody> <tr> <td style=\"text-align:center\">200</td> <td>OK, request proceed without error.</td> </tr> <tr> <td style=\"text-align:center\">204</td> <td>Request proceed successfully and is not returning any content (empty <code>data</code>).</td> </tr> <tr> <td style=\"text-align:center\">400</td> <td>Bad request, either the request header or the API parameter validation failed.</td> </tr> <tr> <td style=\"text-align:center\">401</td> <td>Unauthorized access, either the HTTP request header <code>Authorization</code> is missing or invalid.</td> </tr> <tr> <td style=\"text-align:center\">403</td> <td>Forbidden, the user does not have access to the requested resource.</td> </tr> <tr> <td style=\"text-align:center\">404</td> <td>Not found, requested resource does not exist.</td> </tr> <tr> <td style=\"text-align:center\">429</td> <td>Too many requests sent in a given amount of time, intended for use with rate-limiting schemes.</td> </tr> <tr> <td style=\"text-align:center\">491</td> <td>The API key specified in HTTP request header <code>API-Key</code> is invalid.</td> </tr> <tr> <td style=\"text-align:center\">500</td> <td>Internal server error while processing the request.</td> </tr> </tbody> </table> <h4>API Error Codes</h4> <table> <thead> <tr> <th style=\"text-align:center\"><strong>Code</strong></th> <th><strong>Description</strong></th> </tr> </thead> <tbody> <tr> <td style=\"text-align:center\">40001</td> <td>HTTP request header validation failed.</td> </tr> <tr> <td style=\"text-align:center\">40002</td> <td>API request parameter validation failed.</td> </tr> <tr> <td style=\"text-align:center\">40401</td> <td>The requested resource is not found.</td> </tr> <tr> <td style=\"text-align:center\">49101</td> <td>The API key is not provided.</td> </tr> <tr> <td style=\"text-align:center\">49102</td> <td>Failed to parse the API key, or the API key is invalid.</td> </tr> <tr> <td style=\"text-align:center\">49103</td> <td>The provided API key is not found.</td> </tr> <tr> <td style=\"text-align:center\">49104</td> <td>The provided API key is not intended to be used with the client's app platform.</td> </tr> <tr> <td style=\"text-align:center\">49105</td> <td>The provided API key is not intended to be used with the client's app identifier.</td> </tr> <tr> <td style=\"text-align:center\">49106</td> <td>The API key has expired.</td> </tr> <tr> <td style=\"text-align:center\">49107</td> <td>The API key is disabled.</td> </tr> <tr> <td style=\"text-align:center\">50001</td> <td>An error occurred while validating the API key.</td> </tr> <tr> <td style=\"text-align:center\">99999</td> <td>Other errors, usually without specific reason or action.</td> </tr> </tbody> </table>",
    "error": {
      "fields": {
        "Error 4xx": [
          {
            "group": "Error 4xx",
            "optional": false,
            "field": "HeaderValidationFailed",
            "description": "<p>The request header validation failed.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "HeaderValidationFailed:",
          "content": "HTTP/1.1 200 OK\n{\n  \"status\": 400,\n  \"error\": {\n    \"code\": \"40001\",\n    \"message\": \"Request headers are required (API-Key, Device-Platform)\",\n    \"field\": \"\"\n  },\n  \"data\": {}\n}",
          "type": "json"
        }
      ]
    }
  }
] });
