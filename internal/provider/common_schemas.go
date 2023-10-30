package provider

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"reflect"
	"strings"
	"time"
)

///

func intMinuteTimeout(minute int) *time.Duration {
	return schema.DefaultTimeout(time.Duration(minute) * time.Minute)
}

func schemaTimeouts(create int, read int, update int, del int) *schema.ResourceTimeout {
	return &schema.ResourceTimeout{
		Create: intMinuteTimeout(create),
		Read:   intMinuteTimeout(read),
		Update: intMinuteTimeout(update),
		Delete: intMinuteTimeout(del),
	}
}

func schema10MinuteTimeouts() *schema.ResourceTimeout {
	return schemaTimeouts(10, 10, 10, 10)
}

func schemaTimeoutsCRD(create int, read int, del int) *schema.ResourceTimeout {
	return &schema.ResourceTimeout{
		Create: intMinuteTimeout(create),
		Read:   intMinuteTimeout(read),
		Delete: intMinuteTimeout(del),
	}
}

///

func schemaStringComputedPlain() *schema.Schema {
	return &schema.Schema{Type: schema.TypeString, Computed: true}
}

func schemaStringComputed(description string) *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeString,
		Computed:    true,
		Description: description,
	}
}

func schemaStringRequired(description string) *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Description: description,
	}
}

func schemaStringRequiredNewNotEmpty(description string) *schema.Schema {
	return &schema.Schema{
		Type:         schema.TypeString,
		Required:     true,
		ForceNew:     true,
		ValidateFunc: validation.StringIsNotEmpty,
		Description:  description,
	}
}

func schemaStringRequiredNew(description string) *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		ForceNew:    true,
		Description: description,
	}
}

func schemaStringRequiredNewValidate(description string, validateFunc schema.SchemaValidateFunc) *schema.Schema {
	return &schema.Schema{
		Type:         schema.TypeString,
		Required:     true,
		ForceNew:     true,
		ValidateFunc: validateFunc,
		Description:  description,
	}
}

func schemaStringOptionalComputed(description string) *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		Computed:    true,
		Description: description,
	}
}

func schemaStringOptionalComputedNotEmpty(description string) *schema.Schema {
	return &schema.Schema{
		Type:         schema.TypeString,
		Optional:     true,
		Computed:     true,
		Description:  description,
		ValidateFunc: validation.StringIsNotEmpty,
	}
}

func schemaStringOptionalNew(description string) *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		ForceNew:    true,
		Description: description,
	}
}

func schemaStringOptionalComputedValidate(description string, validateFunc schema.SchemaValidateFunc) *schema.Schema {
	return &schema.Schema{
		Type:         schema.TypeString,
		Optional:     true,
		ValidateFunc: validateFunc,
		Description:  description,
	}
}

func schemaStringOptionalNewValidate(description string, validateFunc schema.SchemaValidateFunc) *schema.Schema {
	return &schema.Schema{
		Type:         schema.TypeString,
		Optional:     true,
		ForceNew:     true,
		ValidateFunc: validateFunc,
		Description:  description,
	}
}

func schemaStringRequiredValidate(description string, validateFunc schema.SchemaValidateFunc) *schema.Schema {
	return &schema.Schema{
		Type:         schema.TypeString,
		Required:     true,
		ValidateFunc: validateFunc,
		Description:  description,
	}
}

func schemaStringRequiredNotEmpty(description string) *schema.Schema {
	return schemaStringRequiredValidate(description, validation.StringIsNotEmpty)
}

func schemaStringOptionalValidateDefault(description string, validateFunc schema.SchemaValidateFunc, defaultValue string) *schema.Schema {
	return &schema.Schema{
		Type:         schema.TypeString,
		Optional:     true,
		Default:      defaultValue,
		ValidateFunc: validateFunc,
		Description:  description,
	}
}

func schemaStringOptionalValidate(description string, validateFunc schema.SchemaValidateFunc) *schema.Schema {
	return &schema.Schema{
		Type:         schema.TypeString,
		Optional:     true,
		ValidateFunc: validateFunc,
		Description:  description,
	}
}

func schemaStringOptionalNotEmpty(description string) *schema.Schema {
	return schemaStringOptionalValidate(description, validation.StringIsNotEmpty)
}

func schemaStringOptionalNewNotEmpty(description string) *schema.Schema {
	return &schema.Schema{
		Type:         schema.TypeString,
		Optional:     true,
		ForceNew:     true,
		Description:  description,
		ValidateFunc: validation.StringIsNotEmpty,
	}
}

///

func schemaStringSetOptional(description string) *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeSet,
		Optional:    true,
		Description: description,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	}
}

func schemaStringListOptional(description string) *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Optional:    true,
		Description: description,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	}
}

func schemaStringListOptionalDescribed(description string, elementDescription string) *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Optional:    true,
		Description: description,
		Elem: &schema.Schema{
			Type:        schema.TypeString,
			Description: elementDescription,
		},
	}
}

func schemaStringListComputedNotEmpty(description string) *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Computed:    true,
		Description: description,
		Elem: &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validation.StringIsNotEmpty,
		},
	}
}

func schemaStringListComputedPlain() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Computed: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	}
}

func schemaStringListRequiredNotEmpty(description string) *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Required:    true,
		Description: description,
		Elem: &schema.Schema{
			Type:         schema.TypeString,
			ValidateFunc: validation.StringIsNotEmpty,
		},
	}
}

func schemaStringListRequiredNewNotEmptyDescribed(description string, elementDescription string) *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Required:    true,
		ForceNew:    true,
		Description: description,
		Elem: &schema.Schema{
			Type:        schema.TypeString,
			Description: elementDescription,
		},
	}
}

///

func schemaStringEnvSensitive(envVariableName string, description string) *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		Sensitive:   true,
		DefaultFunc: schema.EnvDefaultFunc(envVariableName, nil),
		Description: description,
	}
}

func schemaAwsAccountId(description string) *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeString,
		Required:    true,
		ForceNew:    true,
		DefaultFunc: schema.MultiEnvDefaultFunc([]string{PfPfeAwsAccountId, PfeAwsAccountId}, nil),
		Description: description,
	}
}

func schemaAccountUuid(description string) *schema.Schema {
	return &schema.Schema{
		Type:         schema.TypeString,
		Required:     true,
		ForceNew:     true,
		DefaultFunc:  schema.EnvDefaultFunc(PfeAccountId, nil),
		ValidateFunc: validation.IsUUID,
		Description:  description,
	}
}

///

func schemaBoolComputedPlain() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeBool,
		Computed: true,
	}
}

func schemaBoolComputed(description string) *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeBool,
		Computed:    true,
		Description: description,
	}
}

func schemaBoolOptional(description string) *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeBool,
		Optional:    true,
		Description: description,
	}
}

func schemaBoolOptionalNew(description string) *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeBool,
		Optional:    true,
		ForceNew:    true,
		Description: description,
	}
}

func schemaBoolOptionalNewDefault(description string, defaultValue bool) *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeBool,
		Optional:    true,
		ForceNew:    true,
		Default:     defaultValue,
		Description: description,
	}
}

func schemaBoolOptionalDefault(description string, defaultValue bool) *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeBool,
		Optional:    true,
		Default:     defaultValue,
		Description: description,
	}
}

func schemaBoolRequiredNew(description string) *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeBool,
		Required:    true,
		ForceNew:    true,
		Description: description,
	}
}

func schemaBoolRequiredNewDefault(description string, defaultValue bool) *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeBool,
		Optional:    true,
		ForceNew:    true,
		Default:     defaultValue,
		Description: description,
	}
}

///

func schemaIntOptionalNewValidateDefault(description string, validateFunc schema.SchemaValidateFunc, defaultValue int) *schema.Schema {
	return &schema.Schema{
		Type:         schema.TypeInt,
		Optional:     true,
		ForceNew:     true,
		ValidateFunc: validateFunc,
		Default:      defaultValue,
		Description:  description,
	}
}

func schemaIntOptionalNewValidate(description string, validateFunc schema.SchemaValidateFunc) *schema.Schema {
	return &schema.Schema{
		Type:         schema.TypeInt,
		Optional:     true,
		ForceNew:     true,
		ValidateFunc: validateFunc,
		Description:  description,
	}
}

func schemaIntOptionalValidateDefault(description string, validateFunc schema.SchemaValidateFunc, defaultValue int) *schema.Schema {
	return &schema.Schema{
		Type:         schema.TypeInt,
		Optional:     true,
		Default:      defaultValue,
		ValidateFunc: validateFunc,
		Description:  description,
	}
}

func schemaIntOptionalNewDefault(description string, defaultValue int) *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeInt,
		Optional:    true,
		ForceNew:    true,
		Default:     defaultValue,
		Description: description,
	}
}

func schemaIntOptionalValidate(description string, validateFunc schema.SchemaValidateFunc) *schema.Schema {
	return &schema.Schema{
		Type:         schema.TypeInt,
		Optional:     true,
		ValidateFunc: validateFunc,
		Description:  description,
	}
}

func schemaIntOptionalDefault(description string, defaultValue int) *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeInt,
		Optional:    true,
		Default:     defaultValue,
		Description: description,
	}
}

func schemaIntOptional(description string) *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeInt,
		Optional:    true,
		Description: description,
	}
}

func schemaIntComputedPlain() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeInt,
		Computed: true,
	}
}

func schemaIntComputed(description string) *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeInt,
		Computed:    true,
		Description: description,
	}
}

func schemaIntRequired(description string) *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeInt,
		Required:    true,
		Description: description,
	}
}

func schemaIntRequiredNewValidate(description string, validateFunc schema.SchemaValidateFunc) *schema.Schema {
	return &schema.Schema{
		Type:         schema.TypeInt,
		Required:     true,
		ForceNew:     true,
		ValidateFunc: validateFunc,
		Description:  description,
	}
}

func schemaIntRequiredValidate(description string, validateFunc schema.SchemaValidateFunc) *schema.Schema {
	return &schema.Schema{
		Type:         schema.TypeInt,
		Required:     true,
		ValidateFunc: validateFunc,
		Description:  description,
	}
}

///

func schemaFloatComputed(description string) *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeFloat,
		Computed:    true,
		Description: description,
	}
}

///

func stringsToMap(values ...string) map[string]bool {
	result := make(map[string]bool)

	for _, value := range values {
		result[value] = true
	}

	return result
}

func mapStruct(obj interface{}, keys ...string) map[string]interface{} {
	filter := make(map[string]bool)

	for _, key := range keys {
		filter[key] = true
	}

	return structToMap(obj, filter)
}

func structToMap(obj interface{}, filterWords map[string]bool) map[string]interface{} {
	objValue := reflect.ValueOf(obj)

	if objValue.Kind() == reflect.Ptr {
		objValue = objValue.Elem()
	}

	if objValue.Kind() != reflect.Struct {
		return nil
	}

	objType := objValue.Type()
	data := make(map[string]interface{})

	for i := 0; i < objType.NumField(); i++ {
		field := objType.Field(i)
		tag := field.Tag.Get("json")

		// Trim the JSON tag up to the first comma
		if commaIndex := strings.Index(tag, ","); commaIndex != -1 {
			tag = tag[:commaIndex]
		}

		if tag != "" && tag != "-" {
			if 0 == len(filterWords) {
				data[tag] = objValue.Field(i).Interface()
			}
			// Check if the tag exists in the filterWords set
			if _, exists := filterWords[tag]; exists {
				data[tag] = objValue.Field(i).Interface()
			}
		}
	}

	return data
}

func structToMapAll(obj interface{}) map[string]interface{} {
	return structToMap(obj, make(map[string]bool))
}

///

func setResourceDataKeys(d *schema.ResourceData, obj interface{}, keys ...string) error {
	filter := make(map[string]bool)

	for _, key := range keys {
		filter[key] = true
	}

	return setResourceData(d, obj, filter)
}

func setResourceData(d *schema.ResourceData, obj interface{}, filterWords map[string]bool) error {
	objValue := reflect.ValueOf(obj)

	if objValue.Kind() == reflect.Ptr {
		objValue = objValue.Elem()
	}

	if objValue.Kind() != reflect.Struct {
		return fmt.Errorf("unexpected value %v of kind %v", obj, objValue.Kind())
	}

	objType := objValue.Type()

	for i := 0; i < objType.NumField(); i++ {
		field := objType.Field(i)
		tag := field.Tag.Get("json")

		// Trim the JSON tag up to the first comma
		if commaIndex := strings.Index(tag, ","); commaIndex != -1 {
			tag = tag[:commaIndex]
		}

		if tag != "" && tag != "-" {
			if 0 == len(filterWords) {
				_ = d.Set(tag, objValue.Field(i).Interface())
			}
			// Check if the tag exists in the filterWords set
			if _, exists := filterWords[tag]; exists {
				_ = d.Set(tag, objValue.Field(i).Interface())
			}
		}
	}

	return nil
}

//

func schemaPrefixRequired(description string) *schema.Schema {
	return schemaStringRequiredValidate(description, validateIPAddressWithPrefix)
}

func schemaPrefixOptional(description string) *schema.Schema {
	return schemaStringOptionalValidate(description, validateIPAddressWithPrefix)
}
