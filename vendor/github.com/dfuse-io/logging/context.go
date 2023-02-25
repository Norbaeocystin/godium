// Copyright 2019 dfuse Platform Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package logging

import (
	"context"

	"go.uber.org/zap"
)

type loggerKeyType int

const loggerKey loggerKeyType = iota

// WithLogger is used to create a new context with a logger added to it
// so it can be later retrieved using `Logger`.
func WithLogger(ctx context.Context, logger *zap.Logger) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}

// Logger is used to retrieved the logger from the context. If no logger
// is present in the context, the `fallbackLogger` received in parameter
// is returned instead.
func Logger(ctx context.Context, fallbackLogger *zap.Logger) *zap.Logger {
	if ctx == nil {
		return fallbackLogger
	}

	if ctxLogger, ok := ctx.Value(loggerKey).(*zap.Logger); ok {
		return ctxLogger
	}

	return fallbackLogger
}
