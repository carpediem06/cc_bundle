#include <stdio.h>
#include <string.h>
#include <stdlib.h>

static void app_do(void);

/**
 * @brief main
 */
int main(int argc, char *argv[])
{
#ifdef __COVERAGESCANNER__
    __coveragescanner_install("app");
    __coveragescanner_clear();
    __coveragescanner_testname("app_coverage");
#endif
    
    app_do();

#ifdef __COVERAGESCANNER__
   __coveragescanner_teststate("PASSED");
   __coveragescanner_save();
#else
    exit(EXIT_SUCCESS);
#endif
}

/**
 * @brief my app
 */
static void app_do(void)
{
    printf("Hello world\n");
}

/**
 * @brief testcase
 */
void testcase_app() 
{
    printf("\r\n");
    printf("======================\r\n");
    printf("===   UNIT TEST    ===\r\n");
    printf("======================\r\n");
    app_do();
}

#ifdef CU_TEST
/**
 * @brief CU_app
 */
static CU_TestInfo CU_app[] =
{
    {"APP:", testcase_app },
    CU_TEST_INFO_NULL
};
/**
 * @brief scenario
 */
static CU_SuiteInfo scenario[] =
{
  { "APP", NULL, NULL, CU_app },
  CU_SUITE_INFO_NULL
};

/**
 * @brief AddTests
 */
void AddTests(void)
{
    assert(NULL != CU_get_registry());
    assert(!CU_is_test_running());

    CU_register_suites(scenario);

   return;
}
#endif //CU_TEST



