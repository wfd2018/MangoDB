// Copyright 2021 Baltoro OÜ.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package bson

import (
	"testing"
	"time"

	"github.com/MangoDB-io/MangoDB/internal/types"
	"github.com/MangoDB-io/MangoDB/internal/util/testutil"
)

var (
	handshake1 = testCase{
		name: "handshake1",
		v: MustConvertDocument(types.MustMakeDocument(
			"ismaster", true,
			"client", types.MustMakeDocument(
				"driver", types.MustMakeDocument(
					"name", "nodejs",
					"version", "4.0.0-beta.6",
				),
				"os", types.MustMakeDocument(
					"type", "Darwin",
					"name", "darwin",
					"architecture", "x64",
					"version", "20.6.0",
				),
				"platform", "Node.js v14.17.3, LE (unified)|Node.js v14.17.3, LE (unified)",
				"application", types.MustMakeDocument(
					"name", "mongosh 1.0.1",
				),
			),
			"compression", types.Array{"none"},
			"loadBalanced", false,
		)),
		b: testutil.MustParseDumpFile("testdata", "handshake1.hex"),
		j: `{"$k":["ismaster","client","compression","loadBalanced"],"ismaster":true,` +
			`"client":{"$k":["driver","os","platform","application"],"driver":{"$k":["name","version"],` +
			`"name":"nodejs","version":"4.0.0-beta.6"},"os":{"$k":["type","name","architecture","version"],` +
			`"type":"Darwin","name":"darwin","architecture":"x64","version":"20.6.0"},` +
			`"platform":"Node.js v14.17.3, LE (unified)|Node.js v14.17.3, LE (unified)",` +
			`"application":{"$k":["name"],"name":"mongosh 1.0.1"}},"compression":["none"],"loadBalanced":false}`,
	}

	handshake2 = testCase{
		name: "handshake2",
		v: MustConvertDocument(types.MustMakeDocument(
			"ismaster", true,
			"client", types.MustMakeDocument(
				"driver", types.MustMakeDocument(
					"name", "nodejs",
					"version", "4.0.0-beta.6",
				),
				"os", types.MustMakeDocument(
					"type", "Darwin",
					"name", "darwin",
					"architecture", "x64",
					"version", "20.6.0",
				),
				"platform", "Node.js v14.17.3, LE (unified)|Node.js v14.17.3, LE (unified)",
				"application", types.MustMakeDocument(
					"name", "mongosh 1.0.1",
				),
			),
			"compression", types.Array{"none"},
			"loadBalanced", false,
		)),
		b: testutil.MustParseDumpFile("testdata", "handshake2.hex"),
		j: `{"$k":["ismaster","client","compression","loadBalanced"],"ismaster":true,` +
			`"client":{"$k":["driver","os","platform","application"],"driver":{"$k":["name","version"],` +
			`"name":"nodejs","version":"4.0.0-beta.6"},"os":{"$k":["type","name","architecture","version"],` +
			`"type":"Darwin","name":"darwin","architecture":"x64","version":"20.6.0"},` +
			`"platform":"Node.js v14.17.3, LE (unified)|Node.js v14.17.3, LE (unified)",` +
			`"application":{"$k":["name"],"name":"mongosh 1.0.1"}},"compression":["none"],"loadBalanced":false}`,
	}

	handshake3 = testCase{
		name: "handshake3",
		v: MustConvertDocument(types.MustMakeDocument(
			"buildInfo", int32(1),
			"lsid", types.MustMakeDocument(
				"id", types.Binary{
					Subtype: types.BinaryUUID,
					B:       []byte{0xa3, 0x19, 0xf2, 0xb4, 0xa1, 0x75, 0x40, 0xc7, 0xb8, 0xe7, 0xa3, 0xa3, 0x2e, 0xc2, 0x56, 0xbe},
				},
			),
			"$db", "admin",
		)),
		b: testutil.MustParseDumpFile("testdata", "handshake3.hex"),
		j: `{"$k":["buildInfo","lsid","$db"],"buildInfo":1,` +
			`"lsid":{"$k":["id"],"id":{"$b":"oxnytKF1QMe456OjLsJWvg==","s":4}},"$db":"admin"}`,
	}

	handshake4 = testCase{
		name: "handshake4",
		v: MustConvertDocument(types.MustMakeDocument(
			"version", "5.0.0",
			"gitVersion", "1184f004a99660de6f5e745573419bda8a28c0e9",
			"modules", types.Array{},
			"allocator", "tcmalloc",
			"javascriptEngine", "mozjs",
			"sysInfo", "deprecated",
			"versionArray", types.Array{int32(5), int32(0), int32(0), int32(0)},
			"openssl", types.MustMakeDocument(
				"running", "OpenSSL 1.1.1f  31 Mar 2020",
				"compiled", "OpenSSL 1.1.1f  31 Mar 2020",
			),
			"buildEnvironment", types.MustMakeDocument(
				"distmod", "ubuntu2004",
				"distarch", "x86_64",
				"cc", "/opt/mongodbtoolchain/v3/bin/gcc: gcc (GCC) 8.5.0",
				"ccflags", "-Werror -include mongo/platform/basic.h -fasynchronous-unwind-tables -ggdb "+
					"-Wall -Wsign-compare -Wno-unknown-pragmas -Winvalid-pch -fno-omit-frame-pointer "+
					"-fno-strict-aliasing -O2 -march=sandybridge -mtune=generic -mprefer-vector-width=128 "+
					"-Wno-unused-local-typedefs -Wno-unused-function -Wno-deprecated-declarations "+
					"-Wno-unused-const-variable -Wno-unused-but-set-variable -Wno-missing-braces "+
					"-fstack-protector-strong -Wa,--nocompress-debug-sections -fno-builtin-memcmp",
				"cxx", "/opt/mongodbtoolchain/v3/bin/g++: g++ (GCC) 8.5.0",
				"cxxflags", "-Woverloaded-virtual -Wno-maybe-uninitialized -fsized-deallocation -std=c++17",
				"linkflags", "-Wl,--fatal-warnings -pthread -Wl,-z,now -fuse-ld=gold -fstack-protector-strong "+
					"-Wl,--no-threads -Wl,--build-id -Wl,--hash-style=gnu -Wl,-z,noexecstack -Wl,--warn-execstack "+
					"-Wl,-z,relro -Wl,--compress-debug-sections=none -Wl,-z,origin -Wl,--enable-new-dtags",
				"target_arch", "x86_64",
				"target_os", "linux",
				"cppdefines", "SAFEINT_USE_INTRINSICS 0 PCRE_STATIC NDEBUG _XOPEN_SOURCE 700 _GNU_SOURCE "+
					"_REENTRANT 1 _FORTIFY_SOURCE 2 BOOST_THREAD_VERSION 5 BOOST_THREAD_USES_DATETIME "+
					"BOOST_SYSTEM_NO_DEPRECATED BOOST_MATH_NO_LONG_DOUBLE_MATH_FUNCTIONS "+
					"BOOST_ENABLE_ASSERT_DEBUG_HANDLER BOOST_LOG_NO_SHORTHAND_NAMES BOOST_LOG_USE_NATIVE_SYSLOG "+
					"BOOST_LOG_WITHOUT_THREAD_ATTR ABSL_FORCE_ALIGNED_ACCESS",
			),
			"bits", int32(64),
			"debug", false,
			"maxBsonObjectSize", int32(16777216),
			"storageEngines", types.Array{"devnull", "ephemeralForTest", "wiredTiger"},
			"ok", float64(1),
		)),
		b: testutil.MustParseDumpFile("testdata", "handshake4.hex"),
		j: `{"$k":["version","gitVersion","modules","allocator","javascriptEngine","sysInfo","versionArray",` +
			`"openssl","buildEnvironment","bits","debug","maxBsonObjectSize","storageEngines","ok"],` +
			`"version":"5.0.0","gitVersion":"1184f004a99660de6f5e745573419bda8a28c0e9","modules":[],` +
			`"allocator":"tcmalloc","javascriptEngine":"mozjs","sysInfo":"deprecated","versionArray":[5,0,0,0],` +
			`"openssl":{"$k":["running","compiled"],"running":"OpenSSL 1.1.1f  31 Mar 2020",` +
			`"compiled":"OpenSSL 1.1.1f  31 Mar 2020"},` +
			`"buildEnvironment":{"$k":["distmod","distarch","cc","ccflags","cxx","cxxflags","linkflags",` +
			`"target_arch","target_os","cppdefines"],"distmod":"ubuntu2004","distarch":"x86_64",` +
			`"cc":"/opt/mongodbtoolchain/v3/bin/gcc: gcc (GCC) 8.5.0",` +
			`"ccflags":"-Werror -include mongo/platform/basic.h -fasynchronous-unwind-tables -ggdb -Wall ` +
			`-Wsign-compare -Wno-unknown-pragmas -Winvalid-pch -fno-omit-frame-pointer -fno-strict-aliasing ` +
			`-O2 -march=sandybridge -mtune=generic -mprefer-vector-width=128 -Wno-unused-local-typedefs ` +
			`-Wno-unused-function -Wno-deprecated-declarations -Wno-unused-const-variable ` +
			`-Wno-unused-but-set-variable -Wno-missing-braces -fstack-protector-strong ` +
			`-Wa,--nocompress-debug-sections -fno-builtin-memcmp",` +
			`"cxx":"/opt/mongodbtoolchain/v3/bin/g++: g++ (GCC) 8.5.0",` +
			`"cxxflags":"-Woverloaded-virtual -Wno-maybe-uninitialized -fsized-deallocation -std=c++17",` +
			`"linkflags":"-Wl,--fatal-warnings -pthread -Wl,-z,now -fuse-ld=gold -fstack-protector-strong ` +
			`-Wl,--no-threads -Wl,--build-id -Wl,--hash-style=gnu -Wl,-z,noexecstack -Wl,--warn-execstack ` +
			`-Wl,-z,relro -Wl,--compress-debug-sections=none -Wl,-z,origin -Wl,--enable-new-dtags",` +
			`"target_arch":"x86_64","target_os":"linux",` +
			`"cppdefines":"SAFEINT_USE_INTRINSICS 0 PCRE_STATIC NDEBUG _XOPEN_SOURCE 700 _GNU_SOURCE ` +
			`_REENTRANT 1 _FORTIFY_SOURCE 2 BOOST_THREAD_VERSION 5 BOOST_THREAD_USES_DATETIME ` +
			`BOOST_SYSTEM_NO_DEPRECATED BOOST_MATH_NO_LONG_DOUBLE_MATH_FUNCTIONS BOOST_ENABLE_ASSERT_DEBUG_HANDLER ` +
			`BOOST_LOG_NO_SHORTHAND_NAMES BOOST_LOG_USE_NATIVE_SYSLOG BOOST_LOG_WITHOUT_THREAD_ATTR ` +
			`ABSL_FORCE_ALIGNED_ACCESS"},"bits":64,"debug":false,"maxBsonObjectSize":16777216,` +
			`"storageEngines":["devnull","ephemeralForTest","wiredTiger"],"ok":{"$f":"1"}}`,
	}

	all = testCase{
		name: "all",
		v: MustConvertDocument(types.MustMakeDocument(
			"binary", types.Array{
				types.Binary{Subtype: types.BinaryUser, B: []byte{0x42}},
				types.Binary{Subtype: types.BinaryGeneric, B: []byte{}},
			},
			"bool", types.Array{true, false},
			"datetime", types.Array{time.Date(2021, 7, 27, 9, 35, 42, 123000000, time.UTC).Local(), time.Time{}.Local()},
			"double", types.Array{42.13, 0.0},
			"int32", types.Array{int32(42), int32(0)},
			"int64", types.Array{int64(42), int64(0)},
			"objectID", types.Array{types.ObjectID{0x42}, types.ObjectID{}},
			"string", types.Array{"foo", ""},
			"timestamp", types.Array{types.Timestamp(42), types.Timestamp(0)},
		)),
		b: testutil.MustParseDumpFile("testdata", "all.hex"),
		j: `{"$k":["binary","bool","datetime","double","int32","int64","objectID","string","timestamp"],` +
			`"binary":[{"$b":"Qg==","s":128},{"$b":"","s":0}],"bool":[true,false],` +
			`"datetime":[{"$d":"1627378542123"},{"$d":"-62135596800000"}],"double":[{"$f":"42.13"},{"$f":"0"}],` +
			`"int32":[42,0],"int64":[{"$l":"42"},{"$l":"0"}],` +
			`"objectID":[{"$o":"420000000000000000000000"},{"$o":"000000000000000000000000"}],` +
			`"string":["foo",""],"timestamp":[{"$t":"42"},{"$t":"0"}]}`,
	}

	documentTestCases = []testCase{handshake1, handshake2, handshake3, handshake4, all}
)

func TestDocument(t *testing.T) {
	t.Parallel()

	t.Run("Binary", func(t *testing.T) {
		t.Parallel()
		testBinary(t, documentTestCases, func() bsontype { return new(Document) })
	})

	t.Run("JSON", func(t *testing.T) {
		t.Parallel()
		testJSON(t, documentTestCases, func() bsontype { return new(Document) })
	})
}

func FuzzDocumentBinary(f *testing.F) {
	fuzzBinary(f, documentTestCases, func() bsontype { return new(Document) })
}

func FuzzDocumentJSON(f *testing.F) {
	fuzzJSON(f, documentTestCases, func() bsontype { return new(Document) })
}

func BenchmarkDocument(b *testing.B) {
	benchmark(b, documentTestCases, func() bsontype { return new(Document) })
}
