# BrandedFoodObjectItems

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Barcode** | **string** | EAN/UPC barcode | [optional] [default to null]
**Name** | **string** | Item name as provided by brand owner or as shown on packaging | [optional] [default to null]
**Brand** | **string** | The brand name that owns this item | [optional] [default to null]
**Ingredients** | **string** | This food item&#x27;s ingredients from greatest quantity to least | [optional] [default to null]
**Package_** | [***BrandedFoodObjectPackage**](BrandedFoodObject_package.md) |  | [optional] [default to null]
**Serving** | [***BrandedFoodObjectServing**](BrandedFoodObject_serving.md) |  | [optional] [default to null]
**Categories** | **[]string** |  | [optional] [default to null]
**Nutrients** | [**[]BrandedFoodObjectNutrients**](BrandedFoodObject_nutrients.md) | An array containing nutrient informatio objects for this food item | [optional] [default to null]
**DietLabels** | [***BrandedFoodObjectDietLabels**](BrandedFoodObject_diet_labels.md) |  | [optional] [default to null]
**DietFlags** | [**[]BrandedFoodObjectDietFlags**](BrandedFoodObject_diet_flags.md) | An array of ingredient objects that were flagged while grading this item for compatibility with each diet | [optional] [default to null]
**PackagingPhotos** | [***BrandedFoodObjectPackagingPhotos**](BrandedFoodObject_packaging_photos.md) |  | [optional] [default to null]
**Allergens** | **[]string** | An array of ingredients in this item that may cause allergic reactions in people | [optional] [default to null]
**BrandList** | **[]string** | An array of brands we have associated with this item. Some items are sold by more than 1 brand. | [optional] [default to null]
**Countries** | **[]string** | An array of countries where this item is sold | [optional] [default to null]
**CountryDetails** | [***BrandedFoodObjectCountryDetails**](BrandedFoodObject_country_details.md) |  | [optional] [default to null]
**PalmOilIngredients** | **[]string** | An array of ingredients made from palm oil | [optional] [default to null]
**IngredientList** | **[]string** | An array of this item&#x27;s ingredients | [optional] [default to null]
**HasEnglishIngredients** | **bool** | A boolean indicating if we have English ingredients for this item | [optional] [default to null]
**Minerals** | **[]string** | An array of minerals that this item contains | [optional] [default to null]
**Traces** | **[]string** | An array of trace ingredients that may be found in this item | [optional] [default to null]
**Vitamins** | **[]string** | An array of vitamins that are found in this item | [optional] [default to null]
**Description** | **string** | A description of this item | [optional] [default to null]
**Keywords** | **[]string** | An array of keywords that can be used to describe this item | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

