/*
 * Chomp Food Database API Documentation
 *
 * ## Important An **[API key](https://chompthis.com/api/)** is required for access to this API. Get yours at **[https://chompthis.com/api](https://chompthis.com/api/)**.  ### Getting Started   * **[Subscribe](https://chompthis.com/api/#pricing)** to the API.   * Scroll down and click the \"**Authorize**\" button.   * Enter your API key into the \"**value**\" input, click the \"**Authorize**\" button, then click the \"**Close**\" button.   * Scroll down to the section titled \"**default**\" and click on the API endpoint you wish to use.   * Click the \"**Try it out**\" button.   * Enter the information the endpoint requires.   * Click the \"**Execute**\" button.  ### Example    * Branded food response object: **[View example &raquo;](https://raw.githubusercontent.com/chompfoods/examples/master/branded-food-response-object.json)**   * Ingredient response object: **[View example &raquo;](https://raw.githubusercontent.com/chompfoods/examples/master/ingredient-response-object.json)**  ### How Do I Find My API Key?   * Your API key was sent to the email address you used to create your subscription.   * You will also find your API key in the **[Client Center](https://chompthis.com/api/manage.php)**.   * Read **[this article](https://desk.zoho.com/portal/chompthis/kb/articles/how-do-i-find-my-api-key)** for more information.  ### Helpful Links   * **Help & Support**     * [Knowledge Base &raquo;](https://desk.zoho.com/portal/chompthis/kb/chomp)     * [Support &raquo;](https://chompthis.com/api/ticket-new.php)     * [Client Center &raquo;](https://chompthis.com/api/manage.php)   * **Pricing**     * [Subscription Options &raquo;](https://chompthis.com/api/)     * [Cost Calculator &raquo;](https://chompthis.com/api/cost-calculator.php)   * **Guidelines**     * [Terms & License &raquo;](https://chompthis.com/api/terms.php)     * [Attribution &raquo;](https://chompthis.com/api/docs/attribution.php) 
 *
 * API version: 1.0.0-oas3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// An object containing information for this specific item.
type BrandedFoodObjectItems struct {
	// EAN/UPC barcode
	Barcode string `json:"barcode,omitempty"`
	// Item name as provided by brand owner or as shown on packaging
	Name string `json:"name,omitempty"`
	// The brand name that owns this item
	Brand string `json:"brand,omitempty"`
	// This food item's ingredients from greatest quantity to least
	Ingredients string `json:"ingredients,omitempty"`

	Package_ *BrandedFoodObjectPackage `json:"package,omitempty"`

	Serving *BrandedFoodObjectServing `json:"serving,omitempty"`

	Categories []string `json:"categories,omitempty"`
	// An array containing nutrient informatio objects for this food item
	Nutrients []BrandedFoodObjectNutrients `json:"nutrients,omitempty"`

	DietLabels *BrandedFoodObjectDietLabels `json:"diet_labels,omitempty"`
	// An array of ingredient objects that were flagged while grading this item for compatibility with each diet
	DietFlags []BrandedFoodObjectDietFlags `json:"diet_flags,omitempty"`

	PackagingPhotos *BrandedFoodObjectPackagingPhotos `json:"packaging_photos,omitempty"`
	// An array of ingredients in this item that may cause allergic reactions in people
	Allergens []string `json:"allergens,omitempty"`
	// An array of brands we have associated with this item. Some items are sold by more than 1 brand.
	BrandList []string `json:"brand_list,omitempty"`
	// An array of countries where this item is sold
	Countries []string `json:"countries,omitempty"`

	CountryDetails *BrandedFoodObjectCountryDetails `json:"country_details,omitempty"`
	// An array of ingredients made from palm oil
	PalmOilIngredients []string `json:"palm_oil_ingredients,omitempty"`
	// An array of this item's ingredients
	IngredientList []string `json:"ingredient_list,omitempty"`
	// A boolean indicating if we have English ingredients for this item
	HasEnglishIngredients bool `json:"has_english_ingredients,omitempty"`
	// An array of minerals that this item contains
	Minerals []string `json:"minerals,omitempty"`
	// An array of trace ingredients that may be found in this item
	Traces []string `json:"traces,omitempty"`
	// An array of vitamins that are found in this item
	Vitamins []string `json:"vitamins,omitempty"`
	// A description of this item
	Description string `json:"description,omitempty"`
	// An array of keywords that can be used to describe this item
	Keywords []string `json:"keywords,omitempty"`
}
