import { makeModuleRoute } from "./libs/router";

import homeRoutes from './pages/home/routes';
export default [
  makeModuleRoute('/home', homeRoutes),
]
