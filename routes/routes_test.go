package routes

// func Test_routeHandler_getAllSweets(t *testing.T) {
// 	type args struct {
// 		router *mux.Router
// 	}
// 	tests := []struct {
// 		name string
// 		rh   *routeHandler
// 		args args
// 	}{
// 		struct {
// 			name string
// 			rh   *routeHandler
// 			args args
// 		}{
// 			name: "Get all sweets",
// 			rh:   &routeHandler{},
// 			args: args{mux.NewRouter()},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tt.rh.registerRoutes(tt.args.router)
// 			req, _ := http.NewRequest("GET", "/api/sweets", nil)
// 			resp := httptest.NewRecorder()
// 			tt.args.router.ServeHTTP(resp, req)

// 			t.Logf("########## %v", resp.Code)
// 		})
// 	}
// }
