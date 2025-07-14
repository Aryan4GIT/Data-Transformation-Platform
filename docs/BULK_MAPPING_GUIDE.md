# Bulk Mapping Rules Guide

This guide explains how to use the bulk import feature for mapping rules in the Data Mapping API.

## Overview

The bulk import feature allows you to create multiple mapping rules at once using JSON format. This is useful for:
- Migrating existing mapping configurations
- Setting up complex transformations quickly
- Sharing mapping templates between environments

## Using Bulk Import

### 1. Access the Feature
1. Navigate to the Clients page
2. Click on "Mappings" for any client
3. Click the "Bulk Import" button
4. The bulk import dialog will open

### 2. JSON Format

The bulk import expects a JSON array of mapping rule objects. Each object should have:

```json
[
  {
    "source_path": "applicantDetails.0.entityName",
    "destination_path": "applicant_name",
    "transform_type": "copy",
    "transform_logic": "",
    "default_value": "",
    "required": true
  },
  {
    "source_path": "applicantDetails.0.gender",
    "destination_path": "applicant_gender",
    "transform_type": "mapGender",
    "transform_logic": "",
    "default_value": "Unknown",
    "required": false
  }
]
```

### 3. Field Definitions

#### Required Fields:
- `source_path`: Path in the source data (string or array)
- `destination_path`: Path in the output data (string or array)
- `transform_type`: Type of transformation to apply

#### Optional Fields:
- `transform_logic`: Custom logic for expression transforms
- `default_value`: Default value if source is missing
- `required`: Whether the field is required (boolean)

### 4. Path Formats

Paths can be specified as:
- **String format**: `"user.profile.name"` (dot notation)
- **Array format**: `["user", "profile", "name"]`

Both formats are automatically converted to the internal array format.

### 5. Transform Types

Supported transform types:
- `copy`: Direct copy of value
- `toString`: Convert to string
- `toBool`: Convert to boolean
- `toUpperCase`: Convert to uppercase
- `toLowerCase`: Convert to lowercase
- `capitalize`: Capitalize first letter
- `mapGender`: Map gender values (M/MALE → Male, F/FEMALE → Female)
- `formatDate`: Format date values
- `expression`: Custom expression (requires transform_logic)

### 6. Expression Transforms

For expression transforms, provide custom logic:

```json
{
  "source_path": "applicantDetails.0.dob",
  "destination_path": "applicant_age",
  "transform_type": "expression",
  "transform_logic": "2025 - int(split(value, '-')[2])",
  "default_value": "0",
  "required": false
}
```

## Features

### Template Generation
- Click "Load Template" to get a sample JSON structure
- Modify the template with your specific mappings

### Export Existing Rules
- Click "Export Rules" to download current mappings as JSON
- Use exported files as templates for other clients

### Validation
- JSON syntax is validated before import
- Missing required fields are detected
- Invalid transform types are rejected
- Expression syntax is validated

## Example Templates

### Basic Customer Mapping
```json
[
  {
    "source_path": "customer.firstName",
    "destination_path": "first_name",
    "transform_type": "copy",
    "required": true
  },
  {
    "source_path": "customer.lastName",
    "destination_path": "last_name",
    "transform_type": "copy",
    "required": true
  },
  {
    "source_path": "customer.email",
    "destination_path": "email_address",
    "transform_type": "toLowerCase",
    "required": true
  }
]
```

### Financial Data Mapping
```json
[
  {
    "source_path": "loan.amount",
    "destination_path": "principal_amount",
    "transform_type": "toString",
    "required": true
  },
  {
    "source_path": "loan.interestRate",
    "destination_path": "interest_rate",
    "transform_type": "copy",
    "default_value": "0.0",
    "required": false
  },
  {
    "source_path": "loan.status",
    "destination_path": "loan_status",
    "transform_type": "toUpperCase",
    "required": true
  }
]
```

## Best Practices

1. **Start Small**: Test with a few rules before importing large batches
2. **Validate First**: Use the template to ensure correct format
3. **Export Backup**: Export existing rules before bulk importing
4. **Use Descriptive Paths**: Make source and destination paths clear
5. **Set Defaults**: Provide default values for non-required fields
6. **Test Expressions**: Validate custom expressions in the expression help tool

## Troubleshooting

### Common Errors:
- **Invalid JSON**: Check syntax, missing commas, quotes
- **Missing Fields**: Ensure all required fields are present
- **Invalid Transform Type**: Use only supported transform types
- **Expression Errors**: Validate expression syntax

### Tips:
- Use a JSON validator tool for complex imports
- Copy-paste from exported rules for consistency
- Test with a small subset first
- Check the browser console for detailed error messages
