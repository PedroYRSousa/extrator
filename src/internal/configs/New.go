package configs

func New() *S_Configs {
	return &S_Configs{
		PathProducts:                    DEFAULT_PATH_PRODUCTS,
		ProductsParallelExecutionCount:  DEFAULT_PRODUCTS_PARALLEL_EXECUTION_COUNT,
		EndpointsParallelExecutionCount: DEFAULT_ENDPOINTS_PARALLEL_EXECUTION_COUNT,
	}
}
