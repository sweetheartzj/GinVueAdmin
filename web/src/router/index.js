import { createRouter, createWebHistory } from 'vue-router'


const routes = [
    {
        path: "/",
        redirect: "/login",
    },
    {
        path: "/login",
        name: "Login",
        component: () => import("@/views/login/index.vue")
    }
]

const router = createRouter(
    {
        history: createWebHistory(),
        routes
    }
)

export default router
