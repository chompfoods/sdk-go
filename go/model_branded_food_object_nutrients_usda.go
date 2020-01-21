/*
 * Chomp Food Database API Documentation
 *
 * __Important:__   - An __[API key](https://chompthis.com/api/)__ is required for access to this API.   - Get yours at __[https://chompthis.com/api](https://chompthis.com/api/)__.  -----  __Getting Started:__   - __[Subscribe](https://chompthis.com/api/#pricing)__ to the API.   - Scroll down and click the \"__Authorize__\" button.   - Enter your API key into the \"__value__\" input, click the \"__Authorize__\" button, then click the \"__Close__\" button.   - Scroll down to the section titled \"__default__\" and click on the API endpoint you wish to use.   - Click the \"__Try it out__\" button.   - Enter the information the endpoint requires.   - Click the \"__Execute__\" button.  __Example:__    - Branded Food: __[View example](https://raw.githubusercontent.com/chompfoods/examples/master/branded-food-response-object.json)__ API response object.   - Ingredient: __[View example](https://raw.githubusercontent.com/chompfoods/examples/master/ingredient-response-object.json)__ API response object.  -----  __How Do I Find My API Key?__   - Your API key was sent to the email address you used to create your subscription.   - You will also find your API key in the __[Client Center](https://chompthis.com/api/manage.php)__.   - _Read __[this article](https://desk.zoho.com/portal/chompthis/kb/articles/how-do-i-find-my-api-key)__ for more information._  ||| | ------- | -------- | | [Knowledge Base](https://desk.zoho.com/portal/chompthis/kb/chomp) | [Pricing](https://chompthis.com/api/) | | [Attribution](https://chompthis.com/api/docs/attribution.php) | [Cost Calculator](https://chompthis.com/api/cost-calculator.php) | | [Terms & License](https://chompthis.com/api/terms.php) | [Database Search](https://chompthis.com/api/lookup.php) | | [Support](https://chompthis.com/api/ticket-new.php) | [Query Builder](https://chompthis.com/api/build.php) | | [Client Center](https://chompthis.com/api/manage.php) | | 
 *
 * API version: 1.0.0-oas3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type BrandedFoodObjectNutrientsUsda struct {
	// Nutrient ID
	Id int32 `json:"id,omitempty"`
	// Nutrient name
	Name string `json:"name,omitempty"`
	// Amount of the nutrient per 100g of food
	Per100g *BigDecimal `json:"per_100g,omitempty"`
	// The unit used for the measure of this nutrient
	MeasurementUnit string `json:"measurement_unit,omitempty"`
	// Minimum nutrient value
	Min *BigDecimal `json:"min,omitempty"`
	// Maximum nutrient value
	Max *BigDecimal `json:"max,omitempty"`
	// Median nutrient value
	Median *BigDecimal `json:"median,omitempty"`
	// Nutrient rank
	Rank int32 `json:"rank,omitempty"`
	// Number of observations on which the value is based
	DataPoints int32 `json:"data_points,omitempty"`
	// Comments on any unusual aspect of the food nutrient. Examples might include why a nutrient value is different than typically expected.
	Footnote string `json:"footnote,omitempty"`
	// Description of the nutrient source
	Description string `json:"description,omitempty"`
}
