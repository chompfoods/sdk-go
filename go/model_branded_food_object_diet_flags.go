/*
 * Chomp Food Database API Documentation
 *
 * ## Important An **[API key](https://chompthis.com/api/)** is required for access to this API. Get yours at **[https://chompthis.com/api](https://chompthis.com/api/)**.  ### Getting Started   * **[Subscribe](https://chompthis.com/api/#pricing)** to the API.   * Scroll down and click the \"**Authorize**\" button.   * Enter your API key into the \"**value**\" input, click the \"**Authorize**\" button, then click the \"**Close**\" button.   * Scroll down to the section titled \"**default**\" and click on the API endpoint you wish to use.   * Click the \"**Try it out**\" button.   * Enter the information the endpoint requires.   * Click the \"**Execute**\" button.  ### Example    * Branded food response object: **[View example &raquo;](https://raw.githubusercontent.com/chompfoods/examples/master/branded-food-response-object.json)**   * Ingredient response object: **[View example &raquo;](https://raw.githubusercontent.com/chompfoods/examples/master/ingredient-response-object.json)**  ### How Do I Find My API Key?   * Your API key was sent to the email address you used to create your subscription.   * You will also find your API key in the **[Client Center](https://chompthis.com/api/manage.php)**.   * Read **[this article](https://desk.zoho.com/portal/chompthis/kb/articles/how-do-i-find-my-api-key)** for more information.  ||| | ------- | -------- | | [Knowledge Base](https://desk.zoho.com/portal/chompthis/kb/chomp) | [Pricing](https://chompthis.com/api/) | | [Attribution](https://chompthis.com/api/docs/attribution.php) | [Cost Calculator](https://chompthis.com/api/cost-calculator.php) | | [Terms & License](https://chompthis.com/api/terms.php) | [Database Search](https://chompthis.com/api/lookup.php) | | [Support](https://chompthis.com/api/ticket-new.php) | [Query Builder](https://chompthis.com/api/build.php) | | [Client Center](https://chompthis.com/api/manage.php) | | 
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
