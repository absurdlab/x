package healthx

import "go.uber.org/fx"

var Module = fx.Provide(NewHandler)
