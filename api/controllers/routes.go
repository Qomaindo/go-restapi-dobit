package controllers

import "github.com/hdyro/go-restapi-bit/api/middlewares"

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	//================================================================================================================================================================================//
	//================================================================================== API REGION ==================================================================================//
	//================================================================================================================================================================================//
	//COUNTRIES ROUTES
	s.Router.HandleFunc("/web-region/v-0.6/countries", middlewares.SetMiddlewareJSON(s.GetCountries)).Methods("GET")
	s.Router.HandleFunc("/web-region/v-0.6/country/{id}", middlewares.SetMiddlewareJSON(s.GetCountry)).Methods("GET")
	//PROVINCES ROUTES
	s.Router.HandleFunc("/web-region/v-0.6/provinces", middlewares.SetMiddlewareJSON(s.GetProvinces)).Methods("GET")
	s.Router.HandleFunc("/web-region/v-0.6/province/{id}", middlewares.SetMiddlewareJSON(s.GetProvince)).Methods("GET")
	s.Router.HandleFunc("/web-region/v-0.6/province_country/{id}", middlewares.SetMiddlewareJSON(s.GetProvinceCountry)).Methods("GET")
	//REGENCIES ROUTES
	s.Router.HandleFunc("/web-region/v-0.6/regencies", middlewares.SetMiddlewareJSON(s.GetRegencies)).Methods("GET")
	s.Router.HandleFunc("/web-region/v-0.6/regency/{id}", middlewares.SetMiddlewareJSON(s.GetRegency)).Methods("GET")
	s.Router.HandleFunc("/web-region/v-0.6/regency_province/{id}", middlewares.SetMiddlewareJSON(s.GetRegencyProvince)).Methods("GET")
	//DISTRICTS ROUTES
	s.Router.HandleFunc("/web-region/v-0.6/districts", middlewares.SetMiddlewareJSON(s.GetDistricts)).Methods("GET")
	s.Router.HandleFunc("/web-region/v-0.6/district/{id}", middlewares.SetMiddlewareJSON(s.GetDistrict)).Methods("GET")
	s.Router.HandleFunc("/web-region/v-0.6/district_regency/{id}", middlewares.SetMiddlewareJSON(s.GetDistrictRegency)).Methods("GET")
	//VILLAGES ROUTES
	s.Router.HandleFunc("/web-region/v-0.6/villages", middlewares.SetMiddlewareJSON(s.GetVillages)).Methods("GET")
	s.Router.HandleFunc("/web-region/v-0.6/village/{id}", middlewares.SetMiddlewareJSON(s.GetVillage)).Methods("GET")
	s.Router.HandleFunc("/web-region/v-0.6/village_district/{id}", middlewares.SetMiddlewareJSON(s.GetVillageDistrict)).Methods("GET")
	//================================================================================================================================================================================//
	//================================================================================== API SETUP ===================================================================================//
	//================================================================================================================================================================================//
	//ADDRESS ROUTES
	s.Router.HandleFunc("/web/v-0.6/master/address-input-data", middlewares.SetMiddlewareAuthenticationBibite(s.CreateAddressNew)).Methods("POST")
	s.Router.HandleFunc("/web/v-0.6/master/addresss-temp-pendact", middlewares.SetMiddlewareAuthenticationBibite(s.GetAddressTempsPendingActive)).Methods("GET")
	s.Router.HandleFunc("/web/v-0.6/master/address-temp-pendact/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.GetAddressTempPendingActive)).Methods("GET")
	s.Router.HandleFunc("/web/v-0.6/master/address-input-confirm/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.CreateAddressConfirm)).Methods("PUT")
	s.Router.HandleFunc("/web/v-0.6/master/address-input-reject/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.CreateAddressReject)).Methods("DELETE")
	s.Router.HandleFunc("/web/v-0.6/master/addresss-status/{status}", middlewares.SetMiddlewareAuthenticationBibite(s.GetAddresssStatus)).Methods("GET")
	s.Router.HandleFunc("/web/v-0.6/master/address-status/{id}/{status}", middlewares.SetMiddlewareAuthenticationBibite(s.GetAddressStatus)).Methods("GET")
	s.Router.HandleFunc("/web/v-0.6/master/addresss-status-temp/{status}", middlewares.SetMiddlewareAuthenticationBibite(s.GetAddressTempsStatus)).Methods("GET")
	s.Router.HandleFunc("/web/v-0.6/master/address-status-temp/{id}/{status}", middlewares.SetMiddlewareAuthenticationBibite(s.GetAddressTempStatus)).Methods("GET")
	s.Router.HandleFunc("/web/v-0.6/master/address-update-data/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.UpdateAddressNew)).Methods("PUT")
	s.Router.HandleFunc("/web/v-0.6/master/address-update-confirm/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.UpdateAddressConfirm)).Methods("PUT")
	s.Router.HandleFunc("/web/v-0.6/master/address-update-reject/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.UpdateAddressReject)).Methods("PUT")
	s.Router.HandleFunc("/web/v-0.6/master/address-inactive-data/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.InactiveAddressNew)).Methods("PUT")
	s.Router.HandleFunc("/web/v-0.6/master/address-inactive-confirm/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.InactiveAddressConfirm)).Methods("PUT")
	s.Router.HandleFunc("/web/v-0.6/master/address-inactive-reject/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.InactiveAddressReject)).Methods("PUT")
	s.Router.HandleFunc("/web/v-0.6/master/address-active-data/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.ActiveAddressNew)).Methods("PUT")
	s.Router.HandleFunc("/web/v-0.6/master/address-active-confirm/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.ActiveAddressConfirm)).Methods("PUT")
	s.Router.HandleFunc("/web/v-0.6/master/address-active-reject/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.ActiveAddressReject)).Methods("PUT")
	s.Router.HandleFunc("/web/v-0.6/master/address-delete-data/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.DeleteAddressNew)).Methods("PUT")
	s.Router.HandleFunc("/web/v-0.6/master/address-delete-confirm/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.DeleteAddressConfirm)).Methods("PUT")
	s.Router.HandleFunc("/web/v-0.6/master/address-delete-reject/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.DeleteAddressReject)).Methods("PUT")
	s.Router.HandleFunc("/web/v-0.6/master/addresss", middlewares.SetMiddlewareAuthenticationBibite(s.GetAddresss)).Methods("GET")
	s.Router.HandleFunc("/web/v-0.6/master/address/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.GetAddress)).Methods("GET")
	s.Router.HandleFunc("/web/v-0.6/master/addresss-temp", middlewares.SetMiddlewareAuthenticationBibite(s.GetAddressTemps)).Methods("GET")
	s.Router.HandleFunc("/web/v-0.6/master/address-temp/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.GetAddressTemp)).Methods("GET")
	s.Router.HandleFunc("/web/v-0.6/master/address/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.DeleteAddress)).Methods("DELETE")
	s.Router.HandleFunc("/web/v-0.6/master/address-temp/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.DeleteAddressTemp)).Methods("DELETE")

	//================================================================================================================================================================================//
	//===================================================================================== WEB BIBITE ===============================================================================//
	//MOBILE USER
	s.Router.HandleFunc("/web-bibite/v-0.6/mobile/customer-user", middlewares.SetMiddlewareJSON(s.GetMobileUser)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/mobile/customer-user/{id}", middlewares.SetMiddlewareJSON(s.GetDetailUserMobile)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/mobile/customer-user-suspend/{id}", middlewares.SetMiddlewareJSON(s.SuspendMobileUser)).Methods("PUT")
	//================================================================================================================================================================================//
	//BIBITE PRIVILEGE CATEGORY ROUTES
	//================================================================================================================================================================================//
	//BASE CRUD
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-privilege-category", middlewares.SetMiddlewareAuthenticationBibite(s.CreateBibitePrivilegeCategory)).Methods("POST")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-privilege-categorys", middlewares.SetMiddlewareAuthenticationBibite(s.GetBibitePrivilegeCategorysActive)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-privilege-category/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.GetBibitePrivilegeCategoryActive)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-privilege-categorys-delete", middlewares.SetMiddlewareAuthenticationBibite(s.GetBibitePrivilegeCategorysDeleted)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-privilege-category-delete/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.GetBibitePrivilegeCategoryDeleted)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-privilege-category/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.UpdateBibitePrivilegeCategory)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-privilege-category-delete/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.DeleteBibitePrivilegeCategorySoft)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-privilege-category/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.DeleteBibitePrivilegeCategory)).Methods("DELETE")
	//================================================================================================================================================================================//
	//BIBITE PRIVILEGE ROUTES
	//================================================================================================================================================================================//
	//BASE CRUD
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-privilege", middlewares.SetMiddlewareAuthenticationBibite(s.CreateBibitePrivilege)).Methods("POST")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-privileges", middlewares.SetMiddlewareAuthenticationBibite(s.GetBibitePrivilegesActive)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-privilege/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.GetBibitePrivilegeActive)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-privileges-delete", middlewares.SetMiddlewareAuthenticationBibite(s.GetBibitePrivilegesDeleted)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-privilege-delete/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.GetBibitePrivilegeDeleted)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-privilege/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.UpdateBibitePrivilege)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-privilege-delete/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.DeleteBibitePrivilegeSoft)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-privilege/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.DeleteBibitePrivilege)).Methods("DELETE")
	//================================================================================================================================================================================//
	//MITRA PRIVILEGE CATEGORY ROUTES
	//================================================================================================================================================================================//
	//BASE CRUD
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-privilege-category", middlewares.SetMiddlewareAuthenticationBibite(s.CreateMitraPrivilegeCategory)).Methods("POST")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-privilege-categorys", middlewares.SetMiddlewareAuthenticationBibite(s.GetMitraPrivilegeCategorysActive)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-privilege-category/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.GetMitraPrivilegeCategoryActive)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-privilege-categorys-delete", middlewares.SetMiddlewareAuthenticationBibite(s.GetMitraPrivilegeCategorysDeleted)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-privilege-category-delete/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.GetMitraPrivilegeCategoryDeleted)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-privilege-category/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.UpdateMitraPrivilegeCategory)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-privilege-category-delete/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.DeleteMitraPrivilegeCategorySoft)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-privilege-category/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.DeleteMitraPrivilegeCategory)).Methods("DELETE")
	//================================================================================================================================================================================//
	//MITRA PRIVILEGE ROUTES
	//================================================================================================================================================================================//
	//BASE CRUD
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-privilege", middlewares.SetMiddlewareAuthenticationBibite(s.CreateMitraPrivilege)).Methods("POST")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-privileges", middlewares.SetMiddlewareAuthenticationBibite(s.GetMitraPrivilegesActive)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-privilege/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.GetMitraPrivilegeActive)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-privileges-delete", middlewares.SetMiddlewareAuthenticationBibite(s.GetMitraPrivilegesDeleted)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-privilege-delete/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.GetMitraPrivilegeDeleted)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-privilege/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.UpdateMitraPrivilege)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-privilege-delete/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.DeleteMitraPrivilegeSoft)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-privilege/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.DeleteMitraPrivilege)).Methods("DELETE")
	//================================================================================================================================================================================//
	// MITRA TYPE ROUTES
	//================================================================================================================================================================================//
	//BASE CRUD
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-type-input-data", middlewares.SetMiddlewareAuthenticationBibite(s.CreateMitraTypeNew)).Methods("POST")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-type-temp-pendact", middlewares.SetMiddlewareAuthenticationBibite(s.GetMitraTypeTempsPendingActive)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-type-temp-pendact/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.GetMitraTypeTempPendingActive)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-type-input-confirm/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.CreateMitraTypeConfirm)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-type-input-reject/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.CreateMitraTypeReject)).Methods("DELETE")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-types-status/{status}", middlewares.SetMiddlewareAuthenticationBibite(s.GetMitraTypesStatus)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-type-status/{id}/{status}", middlewares.SetMiddlewareAuthenticationBibite(s.GetMitraTypeStatus)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-types-status-temp/{status}", middlewares.SetMiddlewareAuthenticationBibite(s.GetMitraTypeTempsStatus)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-type-status-temp/{id}/{status}", middlewares.SetMiddlewareAuthenticationBibite(s.GetMitraTypeTempStatus)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-type-update-data/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.UpdateMitraTypeNew)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-type-update-confirm/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.UpdateMitraTypeConfirm)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-type-update-reject/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.UpdateMitraTypeReject)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-type-inactive-data/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.InactiveMitraTypeNew)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-type-inactive-confirm/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.InactiveMitraTypeConfirm)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-type-inactive-reject/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.InactiveMitraTypeReject)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-type-active-data/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.ActiveMitraTypeNew)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-type-active-confirm/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.ActiveMitraTypeConfirm)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-type-active-reject/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.ActiveMitraTypeReject)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-type-delete-data/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.DeleteMitraTypeNew)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-type-delete-confirm/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.DeleteMitraTypeConfirm)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-type-delete-reject/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.DeleteMitraTypeReject)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-type", middlewares.SetMiddlewareAuthenticationBibite(s.GetMitraTypes)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-type/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.GetMitraType)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-type-temp", middlewares.SetMiddlewareAuthenticationBibite(s.GetMitraTypeTemps)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-type-temp/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.GetMitraTypeTemp)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-type/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.DeleteMitraType)).Methods("DELETE")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-type-temp/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.DeleteMitraTypeTemp)).Methods("DELETE")
	//================================================================================================================================================================================//
	//PRODUCT DOCUMENT CATEGORY ROUTES
	//================================================================================================================================================================================//
	//BASE CRUD
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-document-category-input-data", middlewares.SetMiddlewareAuthenticationBibite(s.CreateProductDocumentCategoryNew)).Methods("POST")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-documents-category-temp-pendact", middlewares.SetMiddlewareAuthenticationBibite(s.GetProductDocumentCategoryTempsPendingActive)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-document-category-temp-pendact/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.GetProductDocumentCategoryTempPendingActive)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-document-category-input-confirm/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.CreateProductDocumentCategoryConfirm)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-document-category-input-reject/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.CreateProductDocumentCategoryReject)).Methods("DELETE")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-document-categorys-status/{status}", middlewares.SetMiddlewareAuthenticationBibite(s.GetProductDocumentCategorysStatus)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-document-category-status/{id}/{status}", middlewares.SetMiddlewareAuthenticationBibite(s.GetProductDocumentCategoryStatus)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-document-categorys-status-temp/{status}", middlewares.SetMiddlewareAuthenticationBibite(s.GetProductDocumentCategoryTempsStatus)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-document-category-status-temp/{id}/{status}", middlewares.SetMiddlewareAuthenticationBibite(s.GetProductDocumentCategoryTempStatus)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-document-category-update-data/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.UpdateProductDocumentCategoryNew)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-document-category-update-confirm/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.UpdateProductDocumentCategoryConfirm)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-document-category-update-reject/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.UpdateProductDocumentCategoryReject)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-document-category-inactive-data/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.InactiveProductDocumentCategoryNew)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-document-category-inactive-confirm/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.InactiveProductDocumentCategoryConfirm)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-document-category-inactive-reject/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.InactiveProductDocumentCategoryReject)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-document-category-active-data/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.ActiveProductDocumentCategoryNew)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-document-category-active-confirm/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.ActiveProductDocumentCategoryConfirm)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-document-category-active-reject/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.ActiveProductDocumentCategoryReject)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-document-category-delete-data/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.DeleteProductDocumentCategoryNew)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-document-category-delete-confirm/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.DeleteProductDocumentCategoryConfirm)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-document-category-delete-reject/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.DeleteProductDocumentCategoryReject)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-document-category", middlewares.SetMiddlewareAuthenticationBibite(s.GetProductDocumentCategorys)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-document-category/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.GetProductDocumentCategory)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-document-category-temp", middlewares.SetMiddlewareAuthenticationBibite(s.GetProductDocumentCategoryTemps)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-document-category-temp/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.GetProductDocumentCategoryTemp)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-document-category/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.DeleteProductDocumentCategory)).Methods("DELETE")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-document-category-temp/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.DeleteProductDocumentCategoryTemp)).Methods("DELETE")
	//================================================================================================================================================================================//
	//PRODUCT COST ROUTES
	//================================================================================================================================================================================//
	//BASE CRUD
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-cost-input-data", middlewares.SetMiddlewareAuthenticationBibite(s.CreateProductCostNew)).Methods("POST")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-costs-temp-pendact", middlewares.SetMiddlewareAuthenticationBibite(s.GetProductCostTempsPendingActive)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-cost-temp-pendact/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.GetProductCostTempPendingActive)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-cost-input-confirm/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.CreateProductCostConfirm)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-cost-input-reject/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.CreateProductCostReject)).Methods("DELETE")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-costs-status/{status}", middlewares.SetMiddlewareAuthenticationBibite(s.GetProductCostsStatus)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-cost-status/{id}/{status}", middlewares.SetMiddlewareAuthenticationBibite(s.GetProductCostStatus)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-costs-status-temp/{status}", middlewares.SetMiddlewareAuthenticationBibite(s.GetProductCostTempsStatus)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-cost-status-temp/{id}/{status}", middlewares.SetMiddlewareAuthenticationBibite(s.GetProductCostTempStatus)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-cost-update-data/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.UpdateProductCostNew)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-cost-update-confirm/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.UpdateProductCostConfirm)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-cost-update-reject/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.UpdateProductCostReject)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-cost-inactive-data/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.InactiveProductCostNew)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-cost-inactive-confirm/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.InactiveProductCostConfirm)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-cost-inactive-reject/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.InactiveProductCostReject)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-cost-active-data/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.ActiveProductCostNew)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-cost-active-confirm/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.ActiveProductCostConfirm)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-cost-active-reject/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.ActiveProductCostReject)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-cost-delete-data/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.DeleteProductCostNew)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-cost-delete-confirm/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.DeleteProductCostConfirm)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-cost-delete-reject/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.DeleteProductCostReject)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-costs", middlewares.SetMiddlewareAuthenticationBibite(s.GetProductCosts)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-cost/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.GetProductCost)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-costs-temp", middlewares.SetMiddlewareAuthenticationBibite(s.GetProductCostTemps)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-cost-temp/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.GetProductCostTemp)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-cost/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.DeleteProductCost)).Methods("DELETE")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-cost-temp/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.DeleteProductCostTemp)).Methods("DELETE")
	//================================================================================================================================================================================//
	//REFERAL ROUTES
	//================================================================================================================================================================================//
	s.Router.HandleFunc("/web-bibite/v-0.6/master/referal-input-data", middlewares.SetMiddlewareAuthenticationBibite(s.CreateReferalNew)).Methods("POST")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/referals-temp-pendact", middlewares.SetMiddlewareAuthenticationBibite(s.GetReferalTempsPendingActive)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/referal-temp-pendact/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.GetReferalTempPendingActive)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/referal-input-confirm/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.CreateReferalConfirm)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/referal-input-reject/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.CreateReferalReject)).Methods("DELETE")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/referals-status/{status}", middlewares.SetMiddlewareAuthenticationBibite(s.GetReferalsStatus)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/referal-status/{id}/{status}", middlewares.SetMiddlewareAuthenticationBibite(s.GetReferalStatus)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/referals-status-temp/{status}", middlewares.SetMiddlewareAuthenticationBibite(s.GetReferalTempsStatus)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/referal-status-temp/{id}/{status}", middlewares.SetMiddlewareAuthenticationBibite(s.GetReferalTempStatus)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/referal-update-data/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.UpdateReferalNew)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/referal-update-confirm/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.UpdateReferalConfirm)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/referal-update-reject/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.UpdateReferalReject)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/referal-inactive-data/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.InactiveReferalNew)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/referal-inactive-confirm/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.InactiveReferalConfirm)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/referal-inactive-reject/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.InactiveReferalReject)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/referal-active-data/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.ActiveReferalNew)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/referal-active-confirm/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.ActiveReferalConfirm)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/referal-active-reject/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.ActiveReferalReject)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/referal-delete-data/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.DeleteReferalNew)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/referal-delete-confirm/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.DeleteReferalConfirm)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/referal-delete-reject/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.DeleteReferalReject)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/referals", middlewares.SetMiddlewareAuthenticationBibite(s.GetReferals)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/referal/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.GetReferal)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/referals-temp", middlewares.SetMiddlewareAuthenticationBibite(s.GetReferalTemps)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/referal-temp/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.GetReferalTemp)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/referal/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.DeleteReferal)).Methods("DELETE")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/referal-temp/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.DeleteReferalTemp)).Methods("DELETE")
	//================================================================================================================================================================================//
	//COMPANY ROUTES
	//================================================================================================================================================================================//
	s.Router.HandleFunc("/web-bibite/v-0.6/master/company-input-data", middlewares.SetMiddlewareAuthenticationBibite(s.CreateCompanyNew)).Methods("POST")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/companys-temp-pendact", middlewares.SetMiddlewareAuthenticationBibite(s.GetCompanyTempsPendingActive)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/company-temp-pendact/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.GetCompanyTempPendingActive)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/company-input-confirm/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.CreateCompanyConfirm)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/company-input-reject/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.CreateCompanyReject)).Methods("DELETE")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/companys-status/{status}", middlewares.SetMiddlewareAuthenticationBibite(s.GetCompanysStatus)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/company-status/{id}/{status}", middlewares.SetMiddlewareAuthenticationBibite(s.GetCompanyStatus)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/companys-status-temp/{status}", middlewares.SetMiddlewareAuthenticationBibite(s.GetCompanyTempsStatus)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/company-status-temp/{id}/{status}", middlewares.SetMiddlewareAuthenticationBibite(s.GetCompanyTempStatus)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/company-update-data/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.UpdateCompanyNew)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/company-update-confirm/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.UpdateCompanyConfirm)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/company-update-reject/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.UpdateCompanyReject)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/company-inactive-data/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.InactiveCompanyNew)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/company-inactive-confirm/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.InactiveCompanyConfirm)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/company-inactive-reject/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.InactiveCompanyReject)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/company-active-data/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.ActiveCompanyNew)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/company-active-confirm/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.ActiveCompanyConfirm)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/company-active-reject/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.ActiveCompanyReject)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/company-delete-data/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.DeleteCompanyNew)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/company-delete-confirm/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.DeleteCompanyConfirm)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/company-delete-reject/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.DeleteCompanyReject)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/companys", middlewares.SetMiddlewareAuthenticationBibite(s.GetCompanys)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/company/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.GetCompany)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/companys-temp", middlewares.SetMiddlewareAuthenticationBibite(s.GetCompanyTemps)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/company-temp/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.GetCompanyTemp)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/company/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.DeleteCompany)).Methods("DELETE")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/company-temp/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.DeleteCompanyTemp)).Methods("DELETE")
	//================================================================================================================================================================================//
	//PRODUCT DOCUMENT ROUTES
	//================================================================================================================================================================================//
	//BASE CRUD
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-document-input-data", middlewares.SetMiddlewareAuthenticationBibite(s.CreateProductDocumentNew)).Methods("POST")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-documents-temp-pendact", middlewares.SetMiddlewareAuthenticationBibite(s.GetProductDocumentTempsPendingActive)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-document-temp-pendact/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.GetProductDocumentTempPendingActive)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-document-input-confirm/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.CreateProductDocumentConfirm)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-document-input-reject/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.CreateProductDocumentReject)).Methods("DELETE")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-documents-status/{status}", middlewares.SetMiddlewareAuthenticationBibite(s.GetProductDocumentsStatus)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-document-status/{id}/{status}", middlewares.SetMiddlewareAuthenticationBibite(s.GetProductDocumentStatus)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-documents-status-temp/{status}", middlewares.SetMiddlewareAuthenticationBibite(s.GetProductDocumentTempsStatus)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-document-status-temp/{id}/{status}", middlewares.SetMiddlewareAuthenticationBibite(s.GetProductDocumentTempStatus)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-document-update-data/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.UpdateProductDocumentNew)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-document-update-confirm/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.UpdateProductDocumentConfirm)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-document-update-reject/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.UpdateProductDocumentReject)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-document-inactive-data/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.InactiveProductDocumentNew)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-document-inactive-confirm/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.InactiveProductDocumentConfirm)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-document-inactive-reject/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.InactiveProductDocumentReject)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-document-active-data/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.ActiveProductDocumentNew)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-document-active-confirm/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.ActiveProductDocumentConfirm)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-document-active-reject/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.ActiveProductDocumentReject)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-document-delete-data/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.DeleteProductDocumentNew)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-document-delete-confirm/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.DeleteProductDocumentConfirm)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-document-delete-reject/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.DeleteProductDocumentReject)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-documents", middlewares.SetMiddlewareAuthenticationBibite(s.GetProductDocuments)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-document/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.GetProductDocument)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-documents-temp", middlewares.SetMiddlewareAuthenticationBibite(s.GetProductDocumentTemps)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-document-temp/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.GetProductDocumentTemp)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-document/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.DeleteProductDocument)).Methods("DELETE")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/product-document-temp/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.DeleteProductDocumentTemp)).Methods("DELETE")
	//================================================================================================================================================================================//
	//CUSTOMER SUB TYPE ROUTES
	//================================================================================================================================================================================//
	//BASE CRUD
	s.Router.HandleFunc("/web-bibite/v-0.6/master/customer-sub-type-input-data", middlewares.SetMiddlewareAuthenticationBibite(s.CreateCustomerSubTypeNew)).Methods("POST")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/customer-sub-types-temp-pendact", middlewares.SetMiddlewareAuthenticationBibite(s.GetCustomerSubTypeTempsPendingActive)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/customer-sub-type-temp-pendact/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.GetCustomerSubTypeTempPendingActive)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/customer-sub-type-input-confirm/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.CreateCustomerSubTypeConfirm)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/customer-sub-type-input-reject/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.CreateCustomerSubTypeReject)).Methods("DELETE")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/customer-sub-types-status/{status}", middlewares.SetMiddlewareAuthenticationBibite(s.GetCustomerSubTypesStatus)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/customer-sub-type-status/{id}/{status}", middlewares.SetMiddlewareAuthenticationBibite(s.GetCustomerSubTypeStatus)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/customer-sub-types-status-temp/{status}", middlewares.SetMiddlewareAuthenticationBibite(s.GetCustomerSubTypeTempsStatus)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/customer-sub-type-status-temp/{id}/{status}", middlewares.SetMiddlewareAuthenticationBibite(s.GetCustomerSubTypeTempStatus)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/customer-sub-type-update-data/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.UpdateCustomerSubTypeNew)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/customer-sub-type-update-confirm/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.UpdateCustomerSubTypeConfirm)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/customer-sub-type-update-reject/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.UpdateCustomerSubTypeReject)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/customer-sub-type-inactive-data/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.InactiveCustomerSubTypeNew)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/customer-sub-type-inactive-confirm/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.InactiveCustomerSubTypeConfirm)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/customer-sub-type-inactive-reject/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.InactiveCustomerSubTypeReject)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/customer-sub-type-active-data/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.ActiveCustomerSubTypeNew)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/customer-sub-type-active-confirm/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.ActiveCustomerSubTypeConfirm)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/customer-sub-type-active-reject/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.ActiveCustomerSubTypeReject)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/customer-sub-type-delete-data/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.DeleteCustomerSubTypeNew)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/customer-sub-type-delete-confirm/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.DeleteCustomerSubTypeConfirm)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/customer-sub-type-delete-reject/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.DeleteCustomerSubTypeReject)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/customer-sub-types", middlewares.SetMiddlewareAuthenticationBibite(s.GetCustomerSubTypes)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/customer-sub-type/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.GetCustomerSubType)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/customer-sub-types-temp", middlewares.SetMiddlewareAuthenticationBibite(s.GetCustomerSubTypeTemps)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/customer-sub-type-temp/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.GetCustomerSubTypeTemp)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/customer-sub-type/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.DeleteCustomerSubType)).Methods("DELETE")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/customer-sub-type-temp/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.DeleteCustomerSubTypeTemp)).Methods("DELETE")
	//ADDITIONAL ROUTES
	s.Router.HandleFunc("/web-bibite/v-0.6/master/customer-sub-types-by-customer-type/{custtype}", middlewares.SetMiddlewareAuthenticationBibite(s.GetCustomerSubTypesByCustomerType)).Methods("GET")
	//================================================================================================================================================================================//
	//MITRA COMPANY ROUTES
	//================================================================================================================================================================================//
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-company-input-data", middlewares.SetMiddlewareAuthenticationBibite(s.CreateMitraCompanyNew)).Methods("POST")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-companys-temp-pendact", middlewares.SetMiddlewareAuthenticationBibite(s.GetMitraCompanyTempsPendingActive)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-company-temp-pendact/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.GetMitraCompanyTempPendingActive)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-company-input-confirm/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.CreateMitraCompanyConfirm)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-company-input-reject/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.CreateMitraCompanyReject)).Methods("DELETE")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-companys-status/{status}", middlewares.SetMiddlewareAuthenticationBibite(s.GetMitraCompanysStatus)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-company-status/{id}/{status}", middlewares.SetMiddlewareAuthenticationBibite(s.GetMitraCompanyStatus)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-companys-status-temp/{status}", middlewares.SetMiddlewareAuthenticationBibite(s.GetMitraCompanyTempsStatus)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-company-status-temp/{id}/{status}", middlewares.SetMiddlewareAuthenticationBibite(s.GetMitraCompanyTempStatus)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-company-update-data/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.UpdateMitraCompanyNew)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-company-update-confirm/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.UpdateMitraCompanyConfirm)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-company-update-reject/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.UpdateMitraCompanyReject)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-company-inactive-data/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.InactiveMitraCompanyNew)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-company-inactive-confirm/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.InactiveMitraCompanyConfirm)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-company-inactive-reject/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.InactiveMitraCompanyReject)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-company-active-data/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.ActiveMitraCompanyNew)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-company-active-confirm/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.ActiveMitraCompanyConfirm)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-company-active-reject/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.ActiveMitraCompanyReject)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-company-delete-data/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.DeleteMitraCompanyNew)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-company-delete-confirm/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.DeleteMitraCompanyConfirm)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-company-delete-reject/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.DeleteMitraCompanyReject)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-companys", middlewares.SetMiddlewareAuthenticationBibite(s.GetMitraCompanys)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-company/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.GetMitraCompany)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-companys-temp", middlewares.SetMiddlewareAuthenticationBibite(s.GetMitraCompanyTemps)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-company-temp/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.GetMitraCompanyTemp)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-company/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.DeleteMitraCompany)).Methods("DELETE")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-company-temp/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.DeleteMitraCompanyTemp)).Methods("DELETE")
	//================================================================================================================================================================================//
	//BIBITE GROUP ROUTES
	//================================================================================================================================================================================//
	//BASE CRUD
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-group-input-data", middlewares.SetMiddlewareAuthenticationBibite(s.CreateBibiteGroupNew)).Methods("POST")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-groups-status-temp-pendact", middlewares.SetMiddlewareAuthenticationBibite(s.GetBibiteGroupTempsStatusPendingActive)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-group-status-temp-pendact/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.GetBibiteGroupTempStatusPendingActive)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-group-input-confirm/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.CreateBibiteGroupConfirm)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-group-input-reject/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.CreateBibiteGroupReject)).Methods("DELETE")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-groups-status/{status}", middlewares.SetMiddlewareAuthenticationBibite(s.GetBibiteGroupsStatus)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-group-status/{id}/{status}", middlewares.SetMiddlewareAuthenticationBibite(s.GetBibiteGroupStatus)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-groups-status-temp/{status}", middlewares.SetMiddlewareAuthenticationBibite(s.GetBibiteGroupTempsStatus)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-group-status-temp/{id}/{status}", middlewares.SetMiddlewareAuthenticationBibite(s.GetBibiteGroupTempStatus)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-group-update-data/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.UpdateBibiteGroupNew)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-group-update-confirm/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.UpdateBibiteGroupConfirm)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-group-update-reject/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.UpdateBibiteGroupReject)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-group-inactive-data/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.InactiveBibiteGroupNew)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-group-inactive-confirm/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.InactiveBibiteGroupConfirm)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-group-inactive-reject/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.InactiveBibiteGroupReject)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-group-active-data/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.ActiveBibiteGroupNew)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-group-active-confirm/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.ActiveBibiteGroupConfirm)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-group-active-reject/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.ActiveBibiteGroupReject)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-group-delete-data/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.DeleteBibiteGroupNew)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-group-delete-confirm/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.DeleteBibiteGroupConfirm)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-group-delete-reject/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.DeleteBibiteGroupReject)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-groups", middlewares.SetMiddlewareAuthenticationBibite(s.GetBibiteGroups)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-group/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.GetBibiteGroup)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-groups-temp", middlewares.SetMiddlewareAuthenticationBibite(s.GetBibiteGroupTemps)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-group-temp/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.GetBibiteGroupTemp)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-group/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.DeleteBibiteGroup)).Methods("DELETE")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-group-temp/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.DeleteBibiteGroupTemp)).Methods("DELETE")
	//================================================================================================================================================================================//
	//BIBITE ADMINISTRATOR ROUTES
	//================================================================================================================================================================================//
	//BASE CRUD
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-adm-input-data", middlewares.SetMiddlewareAuthenticationBibite(s.CreateBibiteAdministratorNew)).Methods("POST")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-adms-temp-pendact", middlewares.SetMiddlewareAuthenticationBibite(s.GetBibiteAdministratorTempsPendingActive)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-adm-temp-pendact/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.GetBibiteAdministratorTempPendingActive)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-adm-input-confirm/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.CreateBibiteAdministratorConfirm)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-adm-input-reject/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.CreateBibiteAdministratorReject)).Methods("DELETE")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-adms-status/{status}", middlewares.SetMiddlewareAuthenticationBibite(s.GetBibiteAdministratorsStatus)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-adm-status/{id}/{status}", middlewares.SetMiddlewareAuthenticationBibite(s.GetBibiteAdministratorStatus)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-adms-status-temp/{status}", middlewares.SetMiddlewareAuthenticationBibite(s.GetBibiteAdministratorTempsStatus)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-adm-status-temp/{id}/{status}", middlewares.SetMiddlewareAuthenticationBibite(s.GetBibiteAdministratorTempStatus)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-adm-update-data/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.UpdateBibiteAdministratorNew)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-adm-update-confirm/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.UpdateBibiteAdministratorConfirm)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-adm-update-reject/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.UpdateBibiteAdministratorReject)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-adm-inactive-data/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.InactiveBibiteAdministratorNew)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-adm-inactive-confirm/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.InactiveBibiteAdministratorConfirm)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-adm-inactive-reject/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.InactiveBibiteAdministratorReject)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-adm-active-data/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.ActiveBibiteAdministratorNew)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-adm-active-confirm/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.ActiveBibiteAdministratorConfirm)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-adm-active-reject/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.ActiveBibiteAdministratorReject)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-adm-delete-data/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.DeleteBibiteAdministratorNew)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-adm-delete-confirm/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.DeleteBibiteAdministratorConfirm)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-adm-delete-reject/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.DeleteBibiteAdministratorReject)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-adms", middlewares.SetMiddlewareAuthenticationBibite(s.GetBibiteAdministrators)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-adm/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.GetBibiteAdministrator)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-adms-temp", middlewares.SetMiddlewareAuthenticationBibite(s.GetBibiteAdministratorTemps)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-adm-temp/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.GetBibiteAdministratorTemp)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-adm/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.DeleteBibiteAdministrator)).Methods("DELETE")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-adm-temp/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.DeleteBibiteAdministratorTemp)).Methods("DELETE")
	//================================================================================================================================================================================//
	//BIBITE USER ROUTES
	//================================================================================================================================================================================//
	//BASE CRUD
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-user-input-data", middlewares.SetMiddlewareAuthenticationBibite(s.CreateBibiteUserNew)).Methods("POST")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-users-temp-pendact", middlewares.SetMiddlewareAuthenticationBibite(s.GetBibiteUserTempsPendingActive)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-user-temp-pendact/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.GetBibiteUserTempPendingActive)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-user-input-confirm/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.CreateBibiteUserConfirm)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-user-input-reject/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.CreateBibiteUserReject)).Methods("DELETE")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-users-status/{status}", middlewares.SetMiddlewareAuthenticationBibite(s.GetBibiteUsersStatus)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-user-status/{id}/{status}", middlewares.SetMiddlewareAuthenticationBibite(s.GetBibiteUserStatus)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-users-status-temp/{status}", middlewares.SetMiddlewareAuthenticationBibite(s.GetBibiteUserTempsStatus)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-user-status-temp/{id}/{status}", middlewares.SetMiddlewareAuthenticationBibite(s.GetBibiteUserTempStatus)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-user-update-data/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.UpdateBibiteUserNew)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-user-update-confirm/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.UpdateBibiteUserConfirm)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-user-update-reject/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.UpdateBibiteUserReject)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-user-inactive-data/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.InactiveBibiteUserNew)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-user-inactive-confirm/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.InactiveBibiteUserConfirm)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-user-inactive-reject/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.InactiveBibiteUserReject)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-user-active-data/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.ActiveBibiteUserNew)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-user-active-confirm/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.ActiveBibiteUserConfirm)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-user-active-reject/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.ActiveBibiteUserReject)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-user-delete-data/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.DeleteBibiteUserNew)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-user-delete-confirm/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.DeleteBibiteUserConfirm)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-user-delete-reject/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.DeleteBibiteUserReject)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-users", middlewares.SetMiddlewareAuthenticationBibite(s.GetBibiteUsers)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-user/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.GetBibiteUser)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-users-temp", middlewares.SetMiddlewareAuthenticationBibite(s.GetBibiteUserTemps)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-user-temp/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.GetBibiteUserTemp)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-user/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.DeleteBibiteUser)).Methods("DELETE")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/bibite-user-temp/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.DeleteBibiteUserTemp)).Methods("DELETE")
	//================================================================================================================================================================================//
	//MITRA ROUTES
	//================================================================================================================================================================================//
	//BASE CRUD
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-input-data", middlewares.SetMiddlewareAuthenticationBibite(s.CreateMitraNew)).Methods("POST")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitras-temp-pendact", middlewares.SetMiddlewareAuthenticationBibite(s.GetMitraTempsPendingActive)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-temp-pendact/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.GetMitraTempPendingActive)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-input-confirm/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.CreateMitraConfirm)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-input-reject/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.CreateMitraReject)).Methods("DELETE")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitras-status/{status}", middlewares.SetMiddlewareAuthenticationBibite(s.GetMitrasStatus)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-status/{id}/{status}", middlewares.SetMiddlewareAuthenticationBibite(s.GetMitraStatus)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitras-status-temp/{status}", middlewares.SetMiddlewareAuthenticationBibite(s.GetMitraTempsStatus)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-status-temp/{id}/{status}", middlewares.SetMiddlewareAuthenticationBibite(s.GetMitraTempStatus)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-update-data/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.UpdateMitraNew)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-update-confirm/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.UpdateMitraConfirm)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-update-reject/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.UpdateMitraReject)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-inactive-data/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.InactiveMitraNew)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-inactive-confirm/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.InactiveMitraConfirm)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-inactive-reject/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.InactiveMitraReject)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-active-data/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.ActiveMitraNew)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-active-confirm/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.ActiveMitraConfirm)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-active-reject/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.ActiveMitraReject)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-delete-data/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.DeleteMitraNew)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-delete-confirm/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.DeleteMitraConfirm)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-delete-reject/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.DeleteMitraReject)).Methods("PUT")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitras", middlewares.SetMiddlewareAuthenticationBibite(s.GetMitras)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.GetMitra)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitras-temp", middlewares.SetMiddlewareAuthenticationBibite(s.GetMitraTemps)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-temp/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.GetMitraTemp)).Methods("GET")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.DeleteMitra)).Methods("DELETE")
	s.Router.HandleFunc("/web-bibite/v-0.6/master/mitra-temp/{id}", middlewares.SetMiddlewareAuthenticationBibite(s.DeleteMitraTemp)).Methods("DELETE")
	//================================================================================================================================================================================//

	//================================================================================================================================================================================//
	//================================================================================ MOBILE BIBITE =================================================================================//
	//================================================================================================================================================================================//
	// REGISTER ROUTE
	s.Router.HandleFunc("/mobile-bibite/v-0.6/register/savedata", middlewares.SetMiddlewareJSON(s.CustomerRegisterSaveData)).Methods("POST")
	//================================================================================================================================================================================//
	// LOGIN ROUTE
	s.Router.HandleFunc("/mobile-bibite/v-0.6/login/checkphone", middlewares.SetMiddlewareJSON(s.CustomerLoginCheckPhone)).Methods("POST")
	//================================================================================================================================================================================//
	// PIN ROUTE
	s.Router.HandleFunc("/mobile-bibite/v-0.6/pin/set", middlewares.SetMiddlewareAuthenticationCustomer(s.CustomerSetPin)).Methods("PUT")
	s.Router.HandleFunc("/mobile-bibite/v-0.6/pin/verify", middlewares.SetMiddlewareAuthenticationCustomer(s.CustomerVerifyPin)).Methods("POST")
	s.Router.HandleFunc("/mobile-bibite/v-0.6/pin/update", middlewares.SetMiddlewareAuthenticationCustomer(s.CustomerUpdatePin)).Methods("POST")
	//================================================================================================================================================================================//
	// PASSWORD ROUTE
	// s.Router.HandleFunc("/mobile-bibite/v-0.6/password/set/{id}", middlewares.SetMiddlewareJSON(s.CustomerSetPin)).Methods("PUT")
	s.Router.HandleFunc("/mobile-bibite/v-0.6/password/verify", middlewares.SetMiddlewareAuthenticationCustomer(s.CustomerVerifyPassword)).Methods("POST")
	s.Router.HandleFunc("/mobile-bibite/v-0.6/password/update", middlewares.SetMiddlewareAuthenticationCustomer(s.CustomerUpdatePassword)).Methods("POST")
	//================================================================================================================================================================================//
	// PROFILE ROUTE
	s.Router.HandleFunc("/mobile-bibite/v-0.6/customer/dashboard", middlewares.SetMiddlewareAuthenticationCustomer(s.CustomerDashboard)).Methods("POST")
	s.Router.HandleFunc("/mobile-bibite/v-0.6/customer/profile", middlewares.SetMiddlewareAuthenticationCustomer(s.CustomerProfile)).Methods("POST")
	s.Router.HandleFunc("/mobile-bibite/v-0.6/customer/upload-picture", middlewares.SetMiddlewareAuthenticationCustomer(s.uploadProfile)).Methods("POST")
	s.Router.HandleFunc("/mobile-bibite/v-0.6/customer/profile-address", middlewares.SetMiddlewareAuthenticationCustomer(s.CustomerProfileAddress)).Methods("GET")
	s.Router.HandleFunc("/mobile-bibite/v-0.6/customer/profile", middlewares.SetMiddlewareAuthenticationCustomer(s.CustomerUpdateProfile)).Methods("PUT")
	// s.Router.HandleFunc("/mobile-bibite/v-0.6/customer/profile-address-image", middlewares.SetMiddlewareJSON(s.FindProfileAddressImage)).Methods("POST")
	// s.Router.HandleFunc("/mobile-bibite/v-0.6/customer/upload-address-image/{id}", middlewares.SetMiddlewareJSON(s.UploadCustomerProfileAddressImage)).Methods("POST")
	//================================================================================================================================================================================//
	// OTP ROUTE
	s.Router.HandleFunc("/mobile-bibite/v-0.6/otp/send", middlewares.SetMiddlewareAuthenticationCustomer(s.CustomerSendOTP)).Methods("POST")
	s.Router.HandleFunc("/mobile-bibite/v-0.6/otp/verify", middlewares.SetMiddlewareAuthenticationCustomer(s.CustomerVerifyOTP)).Methods("POST")
	//================================================================================================================================================================================//
	// LOAN REQUEST
	s.Router.HandleFunc("/mobile-bibite/v-0.6/customer/checkreferal/{referalcode}", middlewares.SetMiddlewareAuthenticationCustomer(s.GetReferalCheck)).Methods("GET")
	s.Router.HandleFunc("/mobile-bibite/v-0.6/loanreq/mitra", middlewares.SetMiddlewareAuthenticationCustomer(s.GetMitraLoanRequests)).Methods("POST")
	s.Router.HandleFunc("/mobile-bibite/v-0.6/loanreq/create", middlewares.SetMiddlewareAuthenticationCustomer(s.CreateLoanRequest)).Methods("POST")
	s.Router.HandleFunc("/mobile-bibite/v-0.6/upload-collateral-img/{id}", middlewares.SetMiddlewareAuthenticationCustomer(s.UploadCollateralImg)).Methods("POST")
	s.Router.HandleFunc("/mobile-bibite/v-0.6/Get-collateral-img-detail/{id}", middlewares.SetMiddlewareAuthenticationCustomer(s.GetCollateralImg)).Methods("GET")
	s.Router.HandleFunc("/mobile-bibite/v-0.6/loanstatus/list", middlewares.SetMiddlewareAuthenticationCustomer(s.GetCustomerLoanStatusByUserID)).Methods("GET")
	s.Router.HandleFunc("/mobile-bibite/v-0.6/loanstatus/detail/{id}", middlewares.SetMiddlewareAuthenticationCustomer(s.GetCustomerLoanStatusDetailByID)).Methods("GET")
	s.Router.HandleFunc("/mobile-bibite/v-0.6/loanreq/loan-requested/{mitra}", middlewares.SetMiddlewareAuthenticationCustomer(s.GetLoanRequested)).Methods("GET")

	//================================================================================================================================================================================//
	// LOAN OFFERING
	s.Router.HandleFunc("/mobile-bibite/v-0.6/loanoffer/detail/{id}", middlewares.SetMiddlewareAuthenticationCustomer(s.GetCustomerLoanOffering)).Methods("GET")
	s.Router.HandleFunc("/mobile-bibite/v-0.6/loanoffer/accept/{id}", middlewares.SetMiddlewareAuthenticationCustomer(s.UpdateLoanOfferingAccept)).Methods("PUT")
	s.Router.HandleFunc("/mobile-bibite/v-0.6/loanoffer/reject/{id}", middlewares.SetMiddlewareAuthenticationCustomer(s.UpdateLoanOfferingReject)).Methods("PUT")
	//================================================================================================================================================================================//
	// LOAN SCHEDULING
	s.Router.HandleFunc("/mobile-bibite/v-0.6/loanscheduling/set/{id}", middlewares.SetMiddlewareAuthenticationCustomer(s.UpdateCustomerLoanRequestSchedulingSet)).Methods("PUT")
	s.Router.HandleFunc("/mobile-bibite/v-0.6/loanscheduling/reject/{id}", middlewares.SetMiddlewareAuthenticationCustomer(s.RejectLoanRequestScheduling)).Methods("PUT")
	s.Router.HandleFunc("/mobile-bibite/v-0.6/loanscheduling/detail/{id}", middlewares.SetMiddlewareAuthenticationCustomer(s.GetCustomerLoanRequestScheduling)).Methods("GET")
	s.Router.HandleFunc("/mobile-bibite/v-0.6/loanschedulings-by-phone", middlewares.SetMiddlewareAuthenticationCustomer(s.GetAllCustomerLoanRequestSchedulingByPhone)).Methods("GET")
	s.Router.HandleFunc("/mobile-bibite/v-0.6/loanscheduling-by-phone/detail", middlewares.SetMiddlewareAuthenticationCustomer(s.GetCustomerLoanRequestSchedulingByCustomerID)).Methods("GET")
	//================================================================================================================================================================================//
	//================================================================================================================================================================================//
	// GET LOAN MOBILE BY STATUS  9, 10, 11, 12 ROUTE
	s.Router.HandleFunc("/mobile-bibite/v-0.6/get-loans-status/gets-status", middlewares.SetMiddlewareAuthenticationCustomer(s.GetLoansStatus)).Methods("GET")
	//================================================================================================================================================================================//
	// LOAN MOBILE ROUTE
	s.Router.HandleFunc("/mobile-bibite/v-0.6/master/loan-by-id-status/{id}/{status}", middlewares.SetMiddlewareAuthenticationCustomer(s.GetLoanStatus)).Methods("GET")
	//================================================================================================================================================================================//
	// UPDATE CONFIRMATION FUNDS ALREADY TRANSFERED
	s.Router.HandleFunc("/mobile-bibite/v-0.6/loan-confirmation-transfer/update/{id}", middlewares.SetMiddlewareAuthenticationCustomer(s.UpdateLoanConfirmation)).Methods("PUT")
	//================================================================================================================================================================================//
	// GET LOAN REQUEST TIMELINE BIBITE ROUTE
	s.Router.HandleFunc("/mobile-bibite/v-0.6/get-loan-timeline/get-timeline/{id}", middlewares.SetMiddlewareAuthenticationCustomer(s.GetLoanReqTimelineCust)).Methods("GET")

	//================================================================================================================================================================================//

	// LOAN SUBMISSION PRODUCT DOCUMENT IMAGE BIBITE
	s.Router.HandleFunc("/mobile-bibite/v-0.6/upload-document-bibite/{id}", middlewares.SetMiddlewareAuthenticationCustomer(s.CreateLoanSubmissionProdDoctImgCust)).Methods("POST")
	s.Router.HandleFunc("/mobile-bibite/v-0.6/upload-document/get-loan-sbmssn-upload-document-bibite/{id}", middlewares.SetMiddlewareAuthenticationCustomer(s.GetLoanSubmissionProdDoctImg)).Methods("GET")

	s.Router.HandleFunc("/mobile-bibite/v-0.6/get-list-upload-document/{loansubmissionid}", middlewares.SetMiddlewareAuthenticationCustomer(s.GetListUploadDocument)).Methods("GET")
	//================================================================================================================================================================================//

	// NOTIFICATION IN MOBILE
	s.Router.HandleFunc("/mobile-bibite/v-0.6/notification-customer", middlewares.SetMiddlewareAuthenticationCustomer(s.GetNotificationCustomer)).Methods("GET")
	s.Router.HandleFunc("/mobile-bibite/v-0.6/push-notification-bibite-user", middlewares.SetMiddlewareAuthenticationCustomer(s.PushNotif)).Methods("POST")
	s.Router.HandleFunc("/mobile-bibite/v-0.6/bibite-user/read-notif/{id}", middlewares.SetMiddlewareAuthenticationCustomer(s.ReadNotification)).Methods("PUT")
	s.Router.HandleFunc("/mobile-bibite/v-0.6/bibite-user/update-incoming-notif/{id}", middlewares.SetMiddlewareAuthenticationCustomer(s.UpdateNotificationIncoming)).Methods("PUT")
	s.Router.HandleFunc("/mobile-bibite/v-0.6/bibite-user/delete-notif/{id}", middlewares.SetMiddlewareAuthenticationCustomer(s.DeleteNotification)).Methods("PUT")
	//================================================================================================================================================================================//

	//================================================================================================================================================================================//
	//================================================================================ MOBILE MITRA ==================================================================================//
	//================================================================================================================================================================================//
	// LOGIN ROUTE
	s.Router.HandleFunc("/mobile-mitra/v-0.6/login/checkusername", middlewares.SetMiddlewareJSON(s.MitraUserLoginCheckUsername)).Methods("POST")
	s.Router.HandleFunc("/mobile-mitra/v-0.6/login/checkphone", middlewares.SetMiddlewareJSON(s.MitraUserLoginCheckPhone)).Methods("POST")
	//===============================================================================================================================================================================//
	// PIN ROUTE
	s.Router.HandleFunc("/mobile-mitra/v-0.6/pin/set", middlewares.SetMiddlewareAuthenticationMitra(s.MitraUserSetPin)).Methods("PUT")
	s.Router.HandleFunc("/mobile-mitra/v-0.6/pin/verify", middlewares.SetMiddlewareAuthenticationMitra(s.MitraUserVerifyPin)).Methods("POST")
	s.Router.HandleFunc("/mobile-mitra/v-0.6/pin/update", middlewares.SetMiddlewareAuthenticationMitra(s.MitraUserUpdatePin)).Methods("POST")
	//================================================================================================================================================================================//
	// PASSWORD MOBILE ROUTE
	s.Router.HandleFunc("/mobile-mitra/v-0.6/password-mobile/verify", middlewares.SetMiddlewareAuthenticationMitra(s.MitraUserVerifyPassword)).Methods("POST")
	s.Router.HandleFunc("/mobile-mitra/v-0.6/password-mobile/update", middlewares.SetMiddlewareAuthenticationMitra(s.MitraUserUpdatePassword)).Methods("POST")
	//================================================================================================================================================================================//
	// PROFILE ROUTE
	s.Router.HandleFunc("/mobile-mitra/v-0.6/mitra-user/dashboard", middlewares.SetMiddlewareAuthenticationMitra(s.MitraUserDashboard)).Methods("GET")
	s.Router.HandleFunc("/mobile-mitra/v-0.6/mitra-user/profile", middlewares.SetMiddlewareAuthenticationMitra(s.MitraUserProfile)).Methods("GET")
	s.Router.HandleFunc("/mobile-mitra/v-0.6/mitra-user/upload-picture/{id}", middlewares.SetMiddlewareJSON(s.uploadProfileMitra)).Methods("POST")
	s.Router.HandleFunc("/mobile-mitra/v-0.6/mitra-user/profile", middlewares.SetMiddlewareAuthenticationMitra(s.MitraUserUpdateProfile)).Methods("PUT")
	//================================================================================================================================================================================//
	// OTP ROUTE
	s.Router.HandleFunc("/mobile-mitra/v-0.6/otp/send", middlewares.SetMiddlewareAuthenticationMitra(s.MitraUserSendOTP)).Methods("POST")
	s.Router.HandleFunc("/mobile-mitra/v-0.6/otp/verify", middlewares.SetMiddlewareAuthenticationMitra(s.MitraUserVerifyOTP)).Methods("POST")
	//================================================================================================================================================================================//
	// LOAN REQUEST
	s.Router.HandleFunc("/mobile-mitra/v-0.6/loanreq/loan-requested/{mitra}", middlewares.SetMiddlewareAuthenticationMitra(s.GetLoanRequested)).Methods("GET")
	s.Router.HandleFunc("/mobile-mitra/v-0.6/loanreq/loan-requested-detail/{id}", middlewares.SetMiddlewareAuthenticationMitra(s.GetLoanRequestedDetailByID)).Methods("GET")
	// s.Router.HandleFunc("/mobile-mitra/v-0.6/loanreq/loan-requested-detail-kmg/{id}", middlewares.SetMiddlewareAuthenticationMitra(s.GetLoanRequestedKMG)).Methods("GET")
	s.Router.HandleFunc("/mobile-mitra/v-0.6/loanreq/loan-requested-detail-by-mitra-phone/{id}/{phone}", middlewares.SetMiddlewareAuthenticationMitra(s.GetLoanRequestedDetailByIDMitraPhone)).Methods("GET")
	s.Router.HandleFunc("/mobile-mitra/v-0.6/loanreq/loan-requested-accept/{id}", middlewares.SetMiddlewareAuthenticationMitra(s.LoanRequestedAccept)).Methods("PUT")
	s.Router.HandleFunc("/mobile-mitra/v-0.6/loanreq/loan-requested-accepted", middlewares.SetMiddlewareAuthenticationMitra(s.GetLoanRequestedAccepted)).Methods("GET")
	s.Router.HandleFunc("/mobile-mitra/v-0.6/loanreq/loan-requested-count", middlewares.SetMiddlewareAuthenticationMitra(s.GetCountLoanReq)).Methods("GET")
	s.Router.HandleFunc("/mobile-mitra/v-0.6/loanreq/loan-requested-count-proccess", middlewares.SetMiddlewareAuthenticationMitra(s.GetCountLoanReqProccess)).Methods("GET")
	s.Router.HandleFunc("/mobile-mitra/v-0.6/loanreq/loan-requested-accepted-detail/{id}", middlewares.SetMiddlewareAuthenticationMitra(s.GetLoanRequestedAcceptedDetailByID)).Methods("GET")

	//================================================================================================================================================================================//
	// LOAN REQUEST COLLATERAL IMG
	s.Router.HandleFunc("/mobile-mitra/v-0.6/Get-collateral-img-detail/{id}", middlewares.SetMiddlewareJSON(s.GetCollateralImg)).Methods("GET")
	//================================================================================================================================================================================//
	// CUSTOMER SUB TYPE
	s.Router.HandleFunc("/mobile-mitra/v-0.6/get-all/customer-sub-type", middlewares.SetMiddlewareAuthenticationMitra(s.GetCustomerSubTypes)).Methods("GET")

	//================================================================================================================================================================================//
	// LOAN OFFERING
	s.Router.HandleFunc("/mobile-mitra/v-0.6/loanoffer/loan-product-list/{mitra}/{subctgry}", middlewares.SetMiddlewareAuthenticationMitra(s.GetLoanOfferingProduct)).Methods("GET")
	s.Router.HandleFunc("/mobile-mitra/v-0.6/loanoffer/loan-product-detail/{id}", middlewares.SetMiddlewareAuthenticationMitra(s.GetLoanOfferingProductDetail)).Methods("GET")
	s.Router.HandleFunc("/mobile-mitra/v-0.6/loanoffer/loan-product-list-by-mitra/{mitra}", middlewares.SetMiddlewareAuthenticationMitra(s.GetLoanOfferingProductByMitra)).Methods("GET")
	s.Router.HandleFunc("/mobile-mitra/v-0.6/loanoffer/loan-offering-create/{id}", middlewares.SetMiddlewareAuthenticationMitra(s.CreateLoanOffering)).Methods("PUT")
	s.Router.HandleFunc("/mobile-mitra/v-0.6/loanoffer/loan-offering/{phone}", middlewares.SetMiddlewareAuthenticationMitra(s.GetLoanOffered)).Methods("GET")
	s.Router.HandleFunc("/mobile-mitra/v-0.6/loanoffer/loan-offering-detail/{id}", middlewares.SetMiddlewareAuthenticationMitra(s.GetLoanOfferedDetailByID)).Methods("GET")
	s.Router.HandleFunc("/mobile-mitra/v-0.6/loanoffer/detail/{id}", middlewares.SetMiddlewareAuthenticationMitra(s.GetMitraLoanOffering)).Methods("GET")
	//================================================================================================================================================================================//
	// LOAN SCHEDULING
	s.Router.HandleFunc("/mobile-mitra/v-0.6/loanscheduling/detail/{id}", middlewares.SetMiddlewareAuthenticationMitra(s.GetMitraLoanRequestScheduling)).Methods("GET")
	s.Router.HandleFunc("/mobile-mitra/v-0.6/loanschedulings-by-phone", middlewares.SetMiddlewareAuthenticationMitra(s.GetAllMitraLoanRequestSchedulingByPhone)).Methods("GET")
	s.Router.HandleFunc("/mobile-mitra/v-0.6/loanscheduling/set/{id}", middlewares.SetMiddlewareAuthenticationMitra(s.UpdateMitraLoanRequestSchedulingSet)).Methods("PUT")
	//================================================================================================================================================================================//

	// LOAN SUBMISSION
	s.Router.HandleFunc("/mobile-mitra/v-0.6/loansubmission/create", middlewares.SetMiddlewareAuthenticationMitra(s.CreateLoanSubmission)).Methods("POST")
	s.Router.HandleFunc("/mobile-mitra/v-0.6/loansubmission/get-loansubmission-by-req/{id}", middlewares.SetMiddlewareAuthenticationMitra(s.GetLoanSubmissionReqID)).Methods("GET")
	s.Router.HandleFunc("/mobile-mitra/v-0.6/loan-schdl/get-schdl-by-user", middlewares.SetMiddlewareAuthenticationMitra(s.GetLoanMeetSignatureDataByIDUser)).Methods("GET")

	//================================================================================================================================================================================//
	// CREATE T_LOAN
	s.Router.HandleFunc("/mobile-mitra/v-0.6/t-loan/create", middlewares.SetMiddlewareAuthenticationMitra(s.CreateLoan)).Methods("POST")
	s.Router.HandleFunc("/mobile-mitra/v-0.6/t-loan/get-t-loan-by-sbmssn/{id}", middlewares.SetMiddlewareAuthenticationMitra(s.GetLoanBySbmssn)).Methods("GET")
	//================================================================================================================================================================================//
	// Loan SPK ROUTE
	s.Router.HandleFunc("/mobile-mitra/v-0.6/loan-spk/savedata/{id}", middlewares.SetMiddlewareAuthenticationMitra(s.CreateLoanSPK)).Methods("POST")
	s.Router.HandleFunc("/mobile-mitra/v-0.6/loan-spk/gets-loan-spk", middlewares.SetMiddlewareAuthenticationMitra(s.GetLoanSPKs)).Methods("GET")
	s.Router.HandleFunc("/mobile-mitra/v-0.6/loan-spk/get-loan-spk/{id}", middlewares.SetMiddlewareAuthenticationMitra(s.GetLoanSPK)).Methods("GET")
	s.Router.HandleFunc("/mobile-mitra/v-0.6/loan-spk/update/{id}", middlewares.SetMiddlewareAuthenticationMitra(s.UpdateLoanSPK)).Methods("PUT")
	s.Router.HandleFunc("/mobile-mitra/v-0.6/loan-spk/delete/{id}", middlewares.SetMiddlewareAuthenticationMitra(s.DeleteLoanSPK)).Methods("PUT")
	s.Router.HandleFunc("/mobile-mitra/v-0.6/loan-spk/delete-hard/{id}", middlewares.SetMiddlewareAuthenticationMitra(s.DeleteLoanSPK)).Methods("DELETE")

	//================================================================================================================================================================================//
	// Loan Evidence Transfer ROUTE
	s.Router.HandleFunc("/mobile-mitra/v-0.6/loan-evidence-transfer/savedata/{id}", middlewares.SetMiddlewareAuthenticationMitra(s.CreateLoanEvidenceTransfer)).Methods("POST")
	s.Router.HandleFunc("/mobile-mitra/v-0.6/loan-evidence-transfer/gets-loan-evidence-transfer", middlewares.SetMiddlewareAuthenticationMitra(s.GetLoanEvidenceTransfers)).Methods("GET")
	s.Router.HandleFunc("/mobile-mitra/v-0.6/loan-evidence-transfer/get-loan-evidence-transfer/{id}", middlewares.SetMiddlewareAuthenticationMitra(s.GetLoanEvidenceTransfer)).Methods("GET")
	s.Router.HandleFunc("/mobile-mitra/v-0.6/loan-evidence-transfer/update/{id}", middlewares.SetMiddlewareAuthenticationMitra(s.UpdateLoanEvidenceTransfer)).Methods("PUT")
	s.Router.HandleFunc("/mobile-mitra/v-0.6/loan-evidence-transfer/delete/{id}", middlewares.SetMiddlewareAuthenticationMitra(s.DeleteLoanEvidenceTransfer)).Methods("PUT")
	s.Router.HandleFunc("/mobile-mitra/v-0.6/loan-evidence-transfer/delete-hard/{id}", middlewares.SetMiddlewareAuthenticationMitra(s.DeleteHardLoanEvidenceTransfer)).Methods("DELETE")
	//================================================================================================================================================================================//
	// Mitra Product
	s.Router.HandleFunc("/mobile-mitra/v-0.6/mitra-product/get-mitra-product", middlewares.SetMiddlewareAuthenticationMitra(s.GetMitraProductByMitraUserID)).Methods("GET")
	s.Router.HandleFunc("/mobile-mitra/v-0.6/mitra-products/get-mitra-product/{status}", middlewares.SetMiddlewareAuthenticationMitra(s.GetMitraProductsStatus)).Methods("GET")

	//================================================================================================================================================================================//
	// LOAN COMPLETE DOCUMENT
	// s.Router.HandleFunc("/mobile-mitra/v-0.6/loancompletedocument/get-document-list/{id}", middlewares.SetMiddlewareJSON(s.GetLoanCompleteDocumentList)).Methods("GET")
	// s.Router.HandleFunc("/mobile-mitra/v-0.6/loancompletedocument/get-document-attribute/{id}", middlewares.SetMiddlewareJSON(s.GetLoanCompleteDocumentAttribute)).Methods("GET")
	// s.Router.HandleFunc("/mobile-mitra/v-0.6/loancompletedocument/get-document-attribute-value/{id}", middlewares.SetMiddlewareJSON(s.GetLoanCompleteDocumentAttributeValue)).Methods("GET")
	// s.Router.HandleFunc("/mobile-mitra/v-0.6/loancompletedocument/get-document-image/{id}", middlewares.SetMiddlewareJSON(s.GetLoanCompleteDocumentImage)).Methods("GET")
	// s.Router.HandleFunc("/mobile-mitra/v-0.6/loancompletedocument/upload-document-image/{id}", middlewares.SetMiddlewareJSON(s.UploadLoanCompleteDocumentImage)).Methods("PUT")
	// s.Router.HandleFunc("/mobile-mitra/v-0.6/loancompletedocument/update-document-attribute/{id}", middlewares.SetMiddlewareJSON(s.UpdateLoanCompleteDocumentAttribute)).Methods("PUT")
	//================================================================================================================================================================================//
	// LOAN MOBILE ROUTE
	s.Router.HandleFunc("/mobile-mitra/v-0.6/master/loan-by-id-status/{id}/{status}", middlewares.SetMiddlewareAuthenticationMitra(s.GetLoanStatus)).Methods("GET")
	//================================================================================================================================================================================//
	// LOAN MEET COMPLETE MOBILE ROUTE
	s.Router.HandleFunc("/mobile-mitra/v-0.6/loan-meet/create/{id}", middlewares.SetMiddlewareAuthenticationMitra(s.CreateLoanMeet)).Methods("PUT")
	s.Router.HandleFunc("/mobile-mitra/v-0.6/loan-meet/get-loan-meet", middlewares.SetMiddlewareAuthenticationMitra(s.GetLoanMeet)).Methods("GET")
	//================================================================================================================================================================================//
	// LOAN SIGNATURE OF SPK MOBILE ROUTE
	s.Router.HandleFunc("/mobile-mitra/v-0.6/loan-signature/create/{id}", middlewares.SetMiddlewareAuthenticationMitra(s.CreateLoanSignature)).Methods("PUT")
	s.Router.HandleFunc("/mobile-mitra/v-0.6/loan-signature/get-loan-signature/{id}", middlewares.SetMiddlewareAuthenticationMitra(s.GetLoanSignature)).Methods("GET")
	s.Router.HandleFunc("/mobile-mitra/v-0.6/loan-signature-by-id/get-loan-signature/{id}", middlewares.SetMiddlewareAuthenticationMitra(s.GetLoanSignatureByID)).Methods("GET")
	//================================================================================================================================================================================//
	// GET LOAN MOBILE BY STATUS  9, 10, 11, 12 ROUTE
	s.Router.HandleFunc("/mobile-mitra/v-0.6/get-loans-status/gets-status", middlewares.SetMiddlewareAuthenticationMitra(s.GetLoansStatus)).Methods("GET")
	s.Router.HandleFunc("/mobile-mitra/v-0.6/get-data-loans-status/gets-status/{id}", middlewares.SetMiddlewareAuthenticationMitra(s.GetDataLoansStatus)).Methods("GET")
	s.Router.HandleFunc("/mobile-mitra/v-0.6/get-all-data-loans-status/gets-status", middlewares.SetMiddlewareAuthenticationMitra(s.GetAllDataLoansStatus)).Methods("GET")
	//================================================================================================================================================================================//
	// Loan Submission Income Expense
	s.Router.HandleFunc("/mobile-mitra/v-0.6/master/loan-sbmssn-inc-expn-input-data", middlewares.SetMiddlewareAuthenticationMitra(s.CreateLoanSubmissionIncomeExpense)).Methods("POST")
	s.Router.HandleFunc("/mobile-mitra/v-0.6/loan-sbmssn-inc-expn/gets-loan-sbmssn-inc-expn", middlewares.SetMiddlewareAuthenticationMitra(s.GetLoanSubmissionIncomeExpenses)).Methods("GET")
	s.Router.HandleFunc("/mobile-mitra/v-0.6/loan-sbmssn-inc-expn/get-loan-sbmssn-inc-expn/{id}", middlewares.SetMiddlewareAuthenticationMitra(s.GetLoanSubmissionIncomeExpense)).Methods("GET")
	s.Router.HandleFunc("/mobile-mitra/v-0.6/master/loan-sbmssn-inc-expn-update-data/{id}", middlewares.SetMiddlewareAuthenticationMitra(s.UpdateLoanSubmissionIncomeExpense)).Methods("PUT")
	s.Router.HandleFunc("/mobile-mitra/v-0.6/master/loan-sbmssn-inc-expn-delete-data/{id}", middlewares.SetMiddlewareAuthenticationMitra(s.DeleteHardLoanSubmissionIncomeExpense)).Methods("PUT")
	//================================================================================================================================================================================//
	// GET LOAN REQUEST TIMELINE MITRA ROUTE
	s.Router.HandleFunc("/mobile-mitra/v-0.6/loansubmission/create{id}", middlewares.SetMiddlewareAuthenticationMitra(s.GetLoanReqTimelineMitra)).Methods("GET")

	//================================================================================================================================================================================//
	// NOTIFICATION MITRA USER
	// s.Router.HandleFunc("/web-bibite/v-0.6/notification", middlewares.SetMiddlewareJSON(s.testNotif)).Methods("POST")
	s.Router.HandleFunc("/mobile-mitra/v-0.6/notification-mitra-user", middlewares.SetMiddlewareJSON(s.GetNotificationMitraUserByMitraUserID)).Methods("GET")
	s.Router.HandleFunc("/mobile-mitra/v-0.6/push-notification-mitra-user", middlewares.SetMiddlewareJSON(s.PushNotif)).Methods("POST")
	s.Router.HandleFunc("/mobile-mitra/v-0.6/bibite-user/update-incoming-notif/{id}", middlewares.SetMiddlewareJSON(s.UpdateNotificationIncoming)).Methods("PUT")
	s.Router.HandleFunc("/mobile-mitra/v-0.6/mitra-user/read-notif/{id}", middlewares.SetMiddlewareJSON(s.ReadNotification)).Methods("PUT")
	s.Router.HandleFunc("/mobile-mitra/v-0.6/mitra-user/delete-notif/{id}", middlewares.SetMiddlewareJSON(s.DeleteNotification)).Methods("PUT")
	//================================================================================================================================================================================//

	// LOAN SUBMISSION PRODUCT DOCUMENT IMAGE MITRA
	s.Router.HandleFunc("/mobile-mitra/v-0.6/upload-document/{id}", middlewares.SetMiddlewareAuthenticationMitra(s.CreateLoanSubmissionProdDoctImgMitra)).Methods("POST")
	s.Router.HandleFunc("/mobile-mitra/v-0.6/upload-document/get-loan-sbmssn-upload-document/{id}", middlewares.SetMiddlewareAuthenticationMitra(s.GetLoanSubmissionProdDoctImg)).Methods("GET")
	s.Router.HandleFunc("/mobile-mitra/v-0.6/loan-sbmssn-upload-document/{id}", middlewares.SetMiddlewareAuthenticationMitra(s.VerificationLoanSubmissionProdDoctImg)).Methods("PUT")
	s.Router.HandleFunc("/mobile-mitra/v-0.6/update-status-waiting-analys/{id}", middlewares.SetMiddlewareAuthenticationMitra(s.ConfirmDocument)).Methods("PUT")
	//================================================================================================================================================================================//
	//================================================================================================================================================================================//

	// s.Router.HandleFunc("/mobile-mitra/v-0.6/inv/get-inv-code", middlewares.SetMiddlewareJSON(s.Testing)).Methods("GET")

	//================================================================================================================================================================================//
	// API MESSAGE
	s.Router.HandleFunc("/api-message/v-0.6/send-message", middlewares.SetMiddlewareJSON(s.CreateMessage)).Methods("POST")
	s.Router.HandleFunc("/api-message/v-0.6/get-message/get-all-message-by-id/{id}", middlewares.SetMiddlewareJSON(s.GetMessages)).Methods("GET")
	s.Router.HandleFunc("/api-message/v-0.6/get-message/get-detail-message-by-id/{id}", middlewares.SetMiddlewareJSON(s.GetMessageByID)).Methods("GET")
	s.Router.HandleFunc("/api-message/v-0.6/read-message/read-data-message/{id}", middlewares.SetMiddlewareJSON(s.ReadMessage)).Methods("PUT")
	s.Router.HandleFunc("/api-message/v-0.6/delete-message/delete-data-message/{id}", middlewares.SetMiddlewareJSON(s.DeleteMessage)).Methods("PUT")

}
