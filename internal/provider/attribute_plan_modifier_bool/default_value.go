package attribute_plan_modifier_bool

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// defaultValueAttributePlanModifier specifies a default value (types.Bool) for an attribute.
type defaultValueAttributePlanModifier struct {
	DefaultValue types.Bool
}

// DefaultValue is a helper to instantiate a defaultValueAttributePlanModifier.
func DefaultValue(v types.Bool) planmodifier.Bool {
	return &defaultValueAttributePlanModifier{v}
}

var _ planmodifier.Bool = (*defaultValueAttributePlanModifier)(nil)

func (apm *defaultValueAttributePlanModifier) Description(ctx context.Context) string {
	return apm.MarkdownDescription(ctx)
}

func (apm *defaultValueAttributePlanModifier) MarkdownDescription(ctx context.Context) string {
	return fmt.Sprintf("Sets the default value %q (%s) if the attribute is not set", apm.DefaultValue, apm.DefaultValue.Type(ctx))
}

func (apm *defaultValueAttributePlanModifier) PlanModifyBool(_ context.Context, req planmodifier.BoolRequest, res *planmodifier.BoolResponse) {
	// If the attribute configuration is not null, we are done here
	if !req.ConfigValue.IsNull() {
		return
	}

	// If the attribute plan is "known" and "not null", then a previous plan modifier in the sequence
	// has already been applied, and we don't want to interfere.
	if !req.PlanValue.IsUnknown() && !req.PlanValue.IsNull() {
		return
	}

	res.PlanValue = apm.DefaultValue
}
