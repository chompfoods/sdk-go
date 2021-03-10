# {{classname}}

All URIs are relative to *https://chompthis.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FoodBrandedBarcodePhpGet**](DefaultApi.md#FoodBrandedBarcodePhpGet) | **Get** /food/branded/barcode.php | Get a branded food item using a barcode
[**FoodBrandedNamePhpGet**](DefaultApi.md#FoodBrandedNamePhpGet) | **Get** /food/branded/name.php | Get a branded food item by name
[**FoodBrandedSearchPhpGet**](DefaultApi.md#FoodBrandedSearchPhpGet) | **Get** /food/branded/search.php | Get data for branded food items using various search parameters
[**FoodIngredientSearchPhpGet**](DefaultApi.md#FoodIngredientSearchPhpGet) | **Get** /food/ingredient/search.php | Get raw/generic food ingredient item(s)
[**RecipeIdPhpGet**](DefaultApi.md#RecipeIdPhpGet) | **Get** /recipe/id.php | Get a recipe by ID
[**RecipeIngredientPhpGet**](DefaultApi.md#RecipeIngredientPhpGet) | **Get** /recipe/ingredient.php | Get recipes using a list of ingredients
[**RecipeRandomPhpGet**](DefaultApi.md#RecipeRandomPhpGet) | **Get** /recipe/random.php | Get random popular recipes
[**RecipeSearchPhpGet**](DefaultApi.md#RecipeSearchPhpGet) | **Get** /recipe/search.php | Get recipes using a title and optional filters

# **FoodBrandedBarcodePhpGet**
> BrandedFoodObject FoodBrandedBarcodePhpGet(ctx, code, optional)
Get a branded food item using a barcode

## Get data for a branded food using the food's UPC/EAN barcode.  **You must have a Food API key to use this endpoint.** Get a [Food API key](https://chompthis.com/api/).  **Example**  > ```https://chompthis.com/api/v2/food/branded/barcode.php?api_key=API_KEY&code=CODE```  **Tips**   * Read our **[Knowledge Base article](https://desk.zoho.com/portal/chompthis/kb/articles/im-having-trouble-getting-matches-for-barcodes-what-can-id-do)** for helpful tips and tricks.   * Perform a [check-digit](https://en.wikipedia.org/wiki/Check_digit#UPC) on the barcode you are using.   * Use a barcode from our website [ChompThis.com](https://chompthis.com/?r=api). Search for a food and use the barcode shown in the search results.   * It is possible that our database contains the food you're looking for, but does not have the same barcode you are using. This can happen if a manufacturer introduces a variation of the same food, or the barcode you got was from a 2 oz bag of chips when our database has the food packaged in a 4 oz bag.   * [Contact us](https://chompthis.com/contact.php?api=y) if you are having trouble. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **code** | **string**| #### UPC/EAN barcode  **Example** &gt; &#x60;&#x60;&#x60;&amp;code&#x3D;0842234000988&#x60;&#x60;&#x60;  | 
 **optional** | ***DefaultApiFoodBrandedBarcodePhpGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiFoodBrandedBarcodePhpGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **optional.String**| #### **Only required for Premium subscribers.** The unique identifier assigned to each user on your platform. This can be any string of letters or numbers and doesn&#x27;t have to relate to your own database. [Learn more](https://desk.zoho.com/portal/chompthis/en/kb/articles/monthly-active-users)  **Example** &gt; &#x60;&#x60;&#x60;&amp;user_id&#x3D;fehef8w4ha&#x60;&#x60;&#x60;  | 

### Return type

[**BrandedFoodObject**](BrandedFoodObject.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FoodBrandedNamePhpGet**
> BrandedFoodObject FoodBrandedNamePhpGet(ctx, name, optional)
Get a branded food item by name

## Search for branded food items by name.  **You must have a Food API key to use this endpoint.** Get a [Food API key](https://chompthis.com/api/).  **Example** > ```https://chompthis.com/api/v2/food/branded/name.php?api_key=API_KEY&name=NAME```  **Tips**   * Get started by using our **[food lookup tool](https://chompthis.com/api/lookup.php)**.  > This API endpoint is only available to Standard and Premium API subscribers. Please consider upgrading your subscription if you are subscribed to the Limited plan. **[Read this](https://desk.zoho.com/portal/chompthis/kb/articles/can-i-upgrade-downgrade-my-subscription)** if you aren't sure how to upgrade your subscription. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| #### Search for branded food items using a general food name keyword. This does not have to exactly match the \&quot;official\&quot; name for the food.  **Example** &gt; &#x60;&#x60;&#x60;&amp;name&#x3D;Starburst&#x60;&#x60;&#x60;  | 
 **optional** | ***DefaultApiFoodBrandedNamePhpGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiFoodBrandedNamePhpGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| #### Set maximum number of records you want the API to return. The default value is \&quot;**10**.\&quot;  **Example** &gt; &#x60;&#x60;&#x60;&amp;limit&#x3D;10&#x60;&#x60;&#x60;  | 
 **page** | **optional.Int32**| #### This is how you paginate the search result. By default, you will see the first 10 records. You must increment the page number to access the next 10 records, and so on. The default value is \&quot;**1**.\&quot;  **Example** &gt; &#x60;&#x60;&#x60;&amp;page&#x3D;1&#x60;&#x60;&#x60;  | 
 **userId** | **optional.String**| #### **Only required for Premium subscribers.** The unique identifier assigned to each user on your platform. This can be any string of letters or numbers and doesn&#x27;t have to relate to your own database. [Learn more](https://desk.zoho.com/portal/chompthis/en/kb/articles/monthly-active-users)  **Example** &gt; &#x60;&#x60;&#x60;&amp;user_id&#x3D;fehef8w4ha&#x60;&#x60;&#x60;  | 

### Return type

[**BrandedFoodObject**](BrandedFoodObject.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FoodBrandedSearchPhpGet**
> BrandedFoodObject FoodBrandedSearchPhpGet(ctx, optional)
Get data for branded food items using various search parameters

## Search for branded food items using various parameters.  **You must have a Food API key to use this endpoint.** Get a [Food API key](https://chompthis.com/api/).  **Example** > ```https://chompthis.com/api/v2/food/branded/search.php?api_key=API_KEY&brand=BRAND&country=COUNTRY&page=1```  **Tips**    * Get started by using the **[Query Builder](https://chompthis.com/api/build.php)**.  > This API endpoint is only available to Standard and Premium API subscribers. Please consider upgrading your subscription if you are subscribed to the Limited plan. **[Read this](https://desk.zoho.com/portal/chompthis/kb/articles/can-i-upgrade-downgrade-my-subscription)** if you aren't sure how to upgrade your subscription. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiFoodBrandedSearchPhpGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiFoodBrandedSearchPhpGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **allergen** | **optional.String**| #### Filter the search to only include branded foods that contain a specific allergen.  **Example** &gt; &#x60;&#x60;&#x60;&amp;allergen&#x3D;Peanuts&#x60;&#x60;&#x60;  **Important Note**: This parameter cannot be used alone. It must be paired with at least 1 additional parameter.  | 
 **brand** | **optional.String**| #### Filter the search to only include branded foods that are owned by a specific brand.  **Example** &gt; &#x60;&#x60;&#x60;&amp;brand&#x3D;Starbucks&#x60;&#x60;&#x60;  | 
 **category** | **optional.String**| #### Filter the search to only include branded foods from a specific category.  **Example** &gt; &#x60;&#x60;&#x60;&amp;category&#x3D;Plant Based Foods&#x60;&#x60;&#x60;  | 
 **country** | **optional.String**| #### Filter the search to only include branded foods that are sold in a specific country.  **Example** &gt; &#x60;&#x60;&#x60;&amp;country&#x3D;United States&#x60;&#x60;&#x60;  **Important Note**: This parameter cannot be used alone. It must be paired with at least 1 additional parameter.  | 
 **diet** | **optional.String**| #### Filter the search to only include branded foods that are considered compatible with a specific diet.  **Important Note**: This parameter cannot be used alone. It must be paired with at least 1 additional parameter.  | 
 **ingredient** | **optional.String**| #### Filter the search to only include branded foods that contain a specific ingredient.  **Example** &gt; &#x60;&#x60;&#x60;&amp;ingredient&#x3D;Salt&#x60;&#x60;&#x60;  | 
 **keyword** | **optional.String**| #### Filter the search to only include branded foods that are associated with a specific keyword.  **Example** &gt; &#x60;&#x60;&#x60;&amp;keyword&#x3D;Organic&#x60;&#x60;&#x60;  **Important Note**: This parameter cannot be used alone. It must be paired with at least 1 additional parameter.  | 
 **mineral** | **optional.String**| #### Filter the search to only include branded foods that contain a specific mineral.  **Example** &gt; &#x60;&#x60;&#x60;&amp;mineral&#x3D;Potassium&#x60;&#x60;&#x60;  | 
 **nutrient** | **optional.String**| #### Filter the search to only include branded foods that contain a specific nutrient.  **Example** &gt; &#x60;&#x60;&#x60;&amp;nutrient&#x3D;Caffeine&#x60;&#x60;&#x60;  **Important Note**: This parameter cannot be used alone. It must be paired with at least 1 additional parameter.  | 
 **palmOil** | **optional.String**| #### Filter the search to only include branded foods that contain a specific ingredient made using palm oil.  **Example** &gt; &#x60;&#x60;&#x60;&amp;palm_oil&#x3D;E160a Beta Carotene&#x60;&#x60;&#x60;  | 
 **trace** | **optional.String**| ### Filter the search to only include branded foods that contain a specific trace ingredient.  **Example** &gt; &#x60;&#x60;&#x60;&amp;trace&#x3D;Tree Nuts&#x60;&#x60;&#x60;  **Important Note**: This parameter cannot be used alone. It must be paired with at least 1 additional parameter.  | 
 **vitamin** | **optional.String**| #### Filter the search to only include branded foods that contain a specific vitamin.  **Example** &gt; &#x60;&#x60;&#x60;&amp;vitamin&#x3D;Biotin&#x60;&#x60;&#x60;  | 
 **limit** | **optional.Int32**| #### Set maximum number of records you want the API to return. The default value is \&quot;**10**.\&quot;  **Example** &gt; &#x60;&#x60;&#x60;&amp;limit&#x3D;10&#x60;&#x60;&#x60;  | 
 **page** | **optional.Int32**| #### This is how you paginate the search result. By default, you will see the first 10 records. You must increment the page number to access the next 10 records, and so on. The default value is \&quot;**1**.\&quot;  **Example** &gt; &#x60;&#x60;&#x60;&amp;page&#x3D;1&#x60;&#x60;&#x60;  | 
 **userId** | **optional.String**| #### **Only required for Premium subscribers.** The unique identifier assigned to each user on your platform. This can be any string of letters or numbers and doesn&#x27;t have to relate to your own database. [Learn more](https://desk.zoho.com/portal/chompthis/en/kb/articles/monthly-active-users)  **Example** &gt; &#x60;&#x60;&#x60;&amp;user_id&#x3D;fehef8w4ha&#x60;&#x60;&#x60;  | 

### Return type

[**BrandedFoodObject**](BrandedFoodObject.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FoodIngredientSearchPhpGet**
> IngredientObject FoodIngredientSearchPhpGet(ctx, find, optional)
Get raw/generic food ingredient item(s)

## Get data for a specific ingredient or a specific set of ingredients.  **You must have a Food API key to use this endpoint.** Get a [Food API key](https://chompthis.com/api/).  **Example #1: Single Ingredient** > ```https://chompthis.com/api/v2/food/ingredient/search.php?api_key=API_KEY&find=raw broccoli```  **Example #2: Set of Ingredients** > ```https://chompthis.com/api/v2/food/ingredient/search.php?api_key=API_KEY&find=raw broccoli,mashed potatoes,chicken drumstick```  **Tips**   * Expose ingredient endpoints by using our **[food lookup tool](https://chompthis.com/api/lookup.php)**.  > This API endpoint is only available to Standard and Premium API subscribers. Please consider upgrading your subscription if you are subscribed to the Limited plan. **[Read this](https://desk.zoho.com/portal/chompthis/kb/articles/can-i-upgrade-downgrade-my-subscription)** if you aren't sure how to upgrade your subscription. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **find** | **string**| Search our database for a single ingredient or a specific set of ingredients.  **Example #1: Single Ingredient** &gt; &#x60;&#x60;&#x60;&amp;find&#x3D;raw broccoli&#x60;&#x60;&#x60;  **Example #2: Set of Ingredients** &gt; &#x60;&#x60;&#x60;&amp;find&#x3D;raw broccoli,buttermilk waffle,mashed potatoes&#x60;&#x60;&#x60;  **Important Notes**    * Comma-separated lists cannot contain more than **10 ingredients**. You must perform additional API calls if you are looking up more than 10 ingredients.  | 
 **optional** | ***DefaultApiFoodIngredientSearchPhpGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiFoodIngredientSearchPhpGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| #### Set maximum number of records you want the API to return, per search term. The default value is \&quot;**1**.\&quot;  **Example** &gt; &#x60;&#x60;&#x60;&amp;limit&#x3D;3&#x60;&#x60;&#x60;  | 
 **userId** | **optional.String**| #### **Only required for Premium subscribers.** The unique identifier assigned to each user on your platform. This can be any string of letters or numbers and doesn&#x27;t have to relate to your own database. [Learn more](https://desk.zoho.com/portal/chompthis/en/kb/articles/monthly-active-users)  **Example** &gt; &#x60;&#x60;&#x60;&amp;user_id&#x3D;fehef8w4ha&#x60;&#x60;&#x60;  | 

### Return type

[**IngredientObject**](IngredientObject.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RecipeIdPhpGet**
> RecipeObject RecipeIdPhpGet(ctx, id, optional)
Get a recipe by ID

## Get a specific recipe using an ID.  **You must have a Recipe API key to use this endpoint.** Get a [Recipe API key](https://chompthis.com/api/recipes/).  **Example** > ```https://chompthis.com/api/v2/recipe/id.php?api_key=API_KEY&id=RECIPE_ID``` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| #### A recipe ID. Recipe IDs are exposed in the /recipe/search and /recipe/ingredient endpoints.  **Example** &gt; &#x60;&#x60;&#x60;&amp;list&#x3D;tdm_1143_0459d0028fcf8990724785b9e6775037&#x60;&#x60;&#x60;  | 
 **optional** | ***DefaultApiRecipeIdPhpGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiRecipeIdPhpGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userId** | **optional.String**| #### **Only required for Premium subscribers.** The unique identifier assigned to each user on your platform. This can be any string of letters or numbers and doesn&#x27;t have to relate to your own database. [Learn more](https://desk.zoho.com/portal/chompthis/en/kb/articles/monthly-active-users)  **Example** &gt; &#x60;&#x60;&#x60;&amp;user_id&#x3D;fehef8w4ha&#x60;&#x60;&#x60;  | 

### Return type

[**RecipeObject**](RecipeObject.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RecipeIngredientPhpGet**
> RecipeObject RecipeIngredientPhpGet(ctx, list, optional)
Get recipes using a list of ingredients

## Get recipes that include all ingredients from a list.  **You must have a Recipe API key to use this endpoint.** Get a [Recipe API key](https://chompthis.com/api/recipes/).  **Example #1** > ```https://chompthis.com/api/v2/recipe/ingredient.php?api_key=API_KEY&list=INGREDIENT```  **Example #2** > ```https://chompthis.com/api/v2/recipe/ingredient.php?api_key=API_KEY&list=INGREDIENT,INGREDIENT,INGREDIENT``` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **list** | **string**| #### A single ingredient, or a comma-separated list of up to 3 ingredients. Recipes with each of these ingredients will be returned. **You can pass in up to 3 ingredients at a time.**  **Example** &gt; &#x60;&#x60;&#x60;&amp;list&#x3D;cheese,tomato,milk&#x60;&#x60;&#x60;  | 
 **optional** | ***DefaultApiRecipeIngredientPhpGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiRecipeIngredientPhpGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| #### Set maximum number of records you want the API to return. The default value is \&quot;**3**.\&quot;  **Example** &gt; &#x60;&#x60;&#x60;&amp;limit&#x3D;3&#x60;&#x60;&#x60;  | 
 **page** | **optional.Int32**| #### This is how you paginate the search result. By default, you will see the first 3 records. You must increment the page number to access the next 3 records, and so on. The default value is \&quot;**1**.\&quot;  **Example** &gt; &#x60;&#x60;&#x60;&amp;page&#x3D;1&#x60;&#x60;&#x60;  | 
 **userId** | **optional.String**| #### **Only required for Premium subscribers.** The unique identifier assigned to each user on your platform. This can be any string of letters or numbers and doesn&#x27;t have to relate to your own database. [Learn more](https://desk.zoho.com/portal/chompthis/en/kb/articles/monthly-active-users)  **Example** &gt; &#x60;&#x60;&#x60;&amp;user_id&#x3D;fehef8w4ha&#x60;&#x60;&#x60;  | 

### Return type

[**RecipeObject**](RecipeObject.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RecipeRandomPhpGet**
> RecipeObject RecipeRandomPhpGet(ctx, optional)
Get random popular recipes

## Get random recipes that have instructions and nutrients  **You must have a Recipe API key to use this endpoint.** Get a [Recipe API key](https://chompthis.com/api/recipes/).  **Example** > ```https://chompthis.com/api/v2/recipe/random.php?api_key=API_KEY``` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiRecipeRandomPhpGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiRecipeRandomPhpGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| #### Set maximum number of records you want the API to return. The default value is \&quot;**5**.\&quot;  **Example** &gt; &#x60;&#x60;&#x60;&amp;limit&#x3D;5&#x60;&#x60;&#x60;  | 
 **userId** | **optional.String**| #### **Only required for Premium subscribers.** The unique identifier assigned to each user on your platform. This can be any string of letters or numbers and doesn&#x27;t have to relate to your own database. [Learn more](https://desk.zoho.com/portal/chompthis/en/kb/articles/monthly-active-users)  **Example** &gt; &#x60;&#x60;&#x60;&amp;user_id&#x3D;fehef8w4ha&#x60;&#x60;&#x60;  | 

### Return type

[**RecipeObject**](RecipeObject.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RecipeSearchPhpGet**
> RecipeObject RecipeSearchPhpGet(ctx, title, optional)
Get recipes using a title and optional filters

## Get recipes using a title and optional filters.  **You must have a Recipe API key to use this endpoint.** Get a [Recipe API key](https://chompthis.com/api/recipes/).  **Example**  > ```https://chompthis.com/api/v2/recipe/search.php?api_key=API_KEY&title=TITLE``` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **title** | **string**| #### A recipe title  **Example** &gt; &#x60;&#x60;&#x60;&amp;title&#x3D;Banana Bread&#x60;&#x60;&#x60;  | 
 **optional** | ***DefaultApiRecipeSearchPhpGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiRecipeSearchPhpGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **excludedCuisine** | **optional.String**| #### A specific cuisine you want excluded from results  **Example** &gt; &#x60;&#x60;&#x60;&amp;excluded_cuisine&#x3D;Italian&#x60;&#x60;&#x60;  | 
 **includedCuisine** | **optional.String**| #### A specific cuisine you want included in results  **Example** &gt; &#x60;&#x60;&#x60;&amp;included_cuisine&#x3D;Chinese&#x60;&#x60;&#x60;  | 
 **excludedIngredient** | **optional.String**| #### Recipes with this ingredient will be excluded from results  **Example** &gt; &#x60;&#x60;&#x60;&amp;excluded_ingredient&#x3D;egg&#x60;&#x60;&#x60;  | 
 **includedIngredient** | **optional.String**| #### Only recipes with this ingredient will be returned  **Example** &gt; &#x60;&#x60;&#x60;&amp;included_ingredient&#x3D;apple&#x60;&#x60;&#x60;  | 
 **nutrientsRequired** | **optional.Int32**| #### Optionally require all recipes to include nutrition info. Recipes with, or without, nutrition info are returned by default.  **Example** &gt; &#x60;&#x60;&#x60;&amp;nutrients_required&#x3D;1&#x60;&#x60;&#x60;  | 
 **limit** | **optional.Int32**| #### Set maximum number of records you want the API to return. The default value is \&quot;**5**.\&quot;  **Example** &gt; &#x60;&#x60;&#x60;&amp;limit&#x3D;3&#x60;&#x60;&#x60;  | 
 **page** | **optional.Int32**| #### This is how you paginate the search result. By default, you will see the first 5 records. You must increment the page number to access the next 5 records, and so on. The default value is \&quot;**1**.\&quot;  **Example** &gt; &#x60;&#x60;&#x60;&amp;page&#x3D;1&#x60;&#x60;&#x60;  | 
 **userId** | **optional.String**| #### **Only required for Premium subscribers.** The unique identifier assigned to each user on your platform. This can be any string of letters or numbers and doesn&#x27;t have to relate to your own database. [Learn more](https://desk.zoho.com/portal/chompthis/en/kb/articles/monthly-active-users)  **Example** &gt; &#x60;&#x60;&#x60;&amp;user_id&#x3D;fehef8w4ha&#x60;&#x60;&#x60;  | 

### Return type

[**RecipeObject**](RecipeObject.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

