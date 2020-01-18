/*
 * Chomp Food Database API Documentation
 *
 * __Important:__   - An __[API key](https://chompthis.com/api/)__ is required for access to this API.   - Get yours at __[https://chompthis.com/api](https://chompthis.com/api/)__.  -----  __Getting Started:__   - __[Subscribe](https://chompthis.com/api/#pricing)__ to the API.   - Scroll down and click the \"__Authorize__\" button.   - Enter your API key into the \"__value__\" input, click the \"__Authorize__\" button, then click the \"__Close__\" button.   - Scroll down to the section titled \"__default__\" and click on the API endpoint you wish to use.   - Click the \"__Try it out__\" button.   - Enter the information the endpoint requires.   - Click the \"__Execute__\" button.  __Example:__    - __[View example](https://raw.githubusercontent.com/chompfoods/examples/master/response-object.json)__ API response object.  -----  __How Do I Find My API Key?__   - Your API key was sent to the email address you used to create your subscription.   - You will also find your API key in the __[Client Center](https://chompthis.com/api/manage.php)__.   - _Read __[this article](https://desk.zoho.com/portal/chompthis/kb/articles/how-do-i-find-my-api-key)__ for more information._  ||| | ------- | -------- | | [Knowledge Base](https://desk.zoho.com/portal/chompthis/kb/chomp) | [Pricing](https://chompthis.com/api/) | | [Attribution](https://chompthis.com/api/docs/attribution.php) | [Cost Calculator](https://chompthis.com/api/cost-calculator.php) | | [Terms & License](https://chompthis.com/api/terms.php) | [Database Search](https://chompthis.com/api/lookup.php) | | [Support](https://chompthis.com/api/ticket-new.php) | [Query Builder](https://chompthis.com/api/build.php) | | [Client Center](https://chompthis.com/api/manage.php) | | 
 *
 * API version: 1.0.0-oas3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// An object containing information on an individual ingredient that was flagged as potentially not being compatible with a certain diet
type BrandedFoodObjectDietFlags struct {
	// Ingredient name
	Ingredient string `json:"ingredient,omitempty"`
	// Description of the ingredient
	IngredientDescription string `json:"ingredient_description,omitempty"`
	// Name of the diet with which this ingredient may not be compatible
	DietLabel string `json:"diet_label,omitempty"`
	// A description of if we believe this ingredient is compatible with the diet
	IsCompatible string `json:"is_compatible,omitempty"`
	// A numeric representation of if we believe this ingredient is compatible with the diet. Higher values indicate more compatibility
	CompatibilityLevel int32 `json:"compatibility_level,omitempty"`
	// A description of how we graded this ingredient for compatibility with the diet
	CompatibilityDescription string `json:"compatibility_description,omitempty"`
	// Boolean representing if the ingredient is a known allergen
	IsAllergen bool `json:"is_allergen,omitempty"`
}
