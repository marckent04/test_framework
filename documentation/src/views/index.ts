import {
  createWebHistory,
  createRouter,
  type RouteRecordRaw,
} from "vue-router";

import HomeView from "./HomeView.vue";
import SentencesView from "./SentencesView.vue";
import ConfigurationView from "./ConfigurationView.vue";
import QuickStartView from "./QuickStartView.vue";
import GetStartedView from "./GetStartedView.vue";

export const routeNames = {
  home: "home",
  getStarted: "get-started",
  quickstart: "quickstart",
  sentences: "sentences",
  config: "config",
};

const routes: readonly RouteRecordRaw[] = [
  { path: "/", component: HomeView, name: routeNames.home },
  {
    path: "/get-started",
    component: GetStartedView,
    name: routeNames.getStarted,
  },
  {
    path: "/quick-start",
    component: QuickStartView,
    name: routeNames.quickstart,
  },
  { path: "/sentences", component: SentencesView, name: routeNames.sentences },
  {
    path: "/configuration",
    component: ConfigurationView,
    name: routeNames.config,
  },
];

export const router = createRouter({
  history: createWebHistory(),
  routes,
});
