load("@com_google_googleapis//:repository_rules.bzl", "switched_rules_by_language")

googleapis_ext = module_extension(implementation = lambda x: switched_rules_by_language(
    name = "com_google_googleapis_imports",
    cc = True,
))
