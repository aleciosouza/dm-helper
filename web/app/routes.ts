import { index, route, type RouteConfig } from "@react-router/dev/routes";

export default [
    index("routes/index.tsx"),
    route("home", "routes/home.tsx"),
    route("auth", "routes/auth.tsx"),
    route("sheet/:id?", "routes/sheet.tsx"),
] satisfies RouteConfig;
