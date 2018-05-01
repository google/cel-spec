#include "tools/proto_utils_lib.h"

#include "gflags/gflags.h"
#include "google/protobuf/any.pb.h"
#include "google/rpc/status.pb.h"
#include "gmock/gmock.h"
#include "gtest/gtest.h"
#include "absl/strings/str_cat.h"
#include "proto/v1/value.pb.h"
#include "tools/tests/test_message1.pb.h"
#include "tools/tests/test_value.pb.h"

#ifndef CEL_SPEC_OPEN_SOURCE
DECLARE_string(test_srcdir);
#endif

using third_party::cel::spec::tools::tests::TestMessage1;
using third_party::cel::spec::tools::tests::TestValue;

namespace google {
namespace api {
namespace expr {
namespace {

class ProtoUtilsTest : public testing::Test {
 public:
  void SetUp() override {
    std::string main_dir;
#ifdef CEL_SPEC_OPEN_SOURCE
    main_dir = "__main__";
    relative_path_ = "tools/tests/";
#else
    main_dir = "google3";
    relative_path_ = "third_party/cel/spec/tools/tests/";
#endif
    source_path_ = absl::StrCat(std::getenv("TEST_SRCDIR"), "/", main_dir, "/");
  }

  template <typename T>
  void ExpectSuccess(absl::string_view input, absl::string_view message_name,
                     absl::string_view expected) {
    T t;
    google::rpc::Status status;
    std::string output = proto_utils_.TextToBinary(input, message_name, &status);
    EXPECT_EQ(0, status.code());
    t.Clear();
    EXPECT_EQ(true, t.ParseFromString(output));
    EXPECT_EQ(expected, t.name());
  }

  ProtoUtils proto_utils_;
  std::string source_path_;
  std::string relative_path_;
};

TEST_F(ProtoUtilsTest, LinkedinMessage) {
  ExpectSuccess<TestMessage1>("name: 'test1'",
                              "third_party.cel.spec.tools.tests.TestMessage1",
                              "test1");
}

TEST_F(ProtoUtilsTest, LinkedinMessagesAsAny) {
  std::string input = R"(
   name: 'test' value {
      object_value {
        [type.googleapis.com/third_party.cel.spec.tools.tests.TestMessage1]
        {name: 'test1'}}})";
  ExpectSuccess<TestValue>(input, "third_party.cel.spec.tools.tests.TestValue",
                           "test");
}

TEST_F(ProtoUtilsTest, DynamicMessage) {
  google::rpc::Status status = proto_utils_.LoadProtoDefinition(
      absl::StrCat(source_path_, relative_path_, "test_message2.proto"),
      absl::StrCat(relative_path_, "test_message2.proto"));
  EXPECT_EQ(0, status.code());

  proto_utils_.TextToBinary("name2: 'test2'",
                            "third_party.cel.spec.tools.tests.TestMessage2",
                            &status);
  EXPECT_EQ(0, status.code());
}

#ifndef CEL_SPEC_OPEN_SOURCE

TEST_F(ProtoUtilsTest, DynamicMessageAsAny) {
  google::rpc::Status status = proto_utils_.LoadProtoDefinition(
      absl::StrCat(source_path_, relative_path_, "test_message2.proto"),
      absl::StrCat(relative_path_, "test_message2.proto"));
  EXPECT_EQ(0, status.code());

  std::string input = R"(
   name: 'test' value {
      object_value {
        [type.googleapis.com/third_party.cel.spec.tools.tests.TestMessage2]
        {name2: 'test2'}}})";
  ExpectSuccess<TestValue>(input, "third_party.cel.spec.tools.tests.TestValue",
                           "test");
}

#endif

TEST_F(ProtoUtilsTest, DynamicMessageMissingDependency) {
  google::rpc::Status status = proto_utils_.LoadProtoDefinition(
      absl::StrCat(source_path_, relative_path_, "test_message4.proto"),
      absl::StrCat(relative_path_, "test_message4.proto"));
  EXPECT_NE(0, status.code());
}

}  // namespace
}  // namespace expr
}  // namespace api
}  // namespace google
