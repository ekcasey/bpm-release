// Copyright (C) 2017-Present Pivotal Software, Inc. All rights reserved.
//
// This program and the accompanying materials are made available under
// the terms of the under the Apache License, Version 2.0 (the "License”);
// you may not use this file except in compliance with the License.
//
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.  See the
// License for the specific language governing permissions and limitations
// under the License.

package bpm_test

import (
	"bpm/bpm"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Config", func() {
	Describe("ParseConfig", func() {
		var configPath string

		BeforeEach(func() {
			configPath = "fixtures/example.yml"
		})

		It("parses a yaml file into a bpm config", func() {
			cfg, err := bpm.ParseConfig(configPath)
			Expect(err).NotTo(HaveOccurred())

			expectedMemoryLimit := "100G"
			expectedOpenFilesLimit := uint64(100)
			Expect(cfg.Executable).To(Equal("/var/vcap/packages/program/bin/program-server"))
			Expect(cfg.Args).To(ConsistOf("--port=2424", "--host=\"localhost\""))
			Expect(cfg.Env).To(ConsistOf("FOO=BAR", "BAZ=BUZZ"))
			Expect(cfg.Limits.Memory).To(Equal(&expectedMemoryLimit))
			Expect(cfg.Limits.OpenFiles).To(Equal(&expectedOpenFilesLimit))
		})

		Context("when reading the file fails", func() {
			BeforeEach(func() {
				configPath = "does-not-exist"
			})

			It("returns an error", func() {
				_, err := bpm.ParseConfig(configPath)
				Expect(err).To(HaveOccurred())
			})
		})

		Context("when the yaml is invalid", func() {
			BeforeEach(func() {
				configPath = "fixtures/example-invalid-yaml.yml"
			})

			It("returns an error", func() {
				_, err := bpm.ParseConfig(configPath)
				Expect(err).To(HaveOccurred())
			})
		})

		Context("when the configuration is not valid", func() {
			BeforeEach(func() {
				configPath = "fixtures/example-invalid.yml"
			})

			It("returns an error", func() {
				_, err := bpm.ParseConfig(configPath)
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("Validate", func() {
		var cfg *bpm.Config

		BeforeEach(func() {
			cfg = &bpm.Config{
				Executable: "executable",
			}
		})

		It("does not error on a valid config", func() {
			Expect(cfg.Validate()).To(Succeed())
		})

		Context("when the config does not have an Executable", func() {
			BeforeEach(func() {
				cfg.Executable = ""
			})

			It("returns an error", func() {
				Expect(cfg.Validate()).To(HaveOccurred())
			})
		})
	})
})
