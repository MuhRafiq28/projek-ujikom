import Register from '@/views/Register.vue';

const routes = [
  { path: '/register', component: Register },
  { path: '/', component: IndexPage },
  { path: '/home', component: HomePage },
  { path: '/login', component: LoginPage }
];

export default new VueRouter({ routes });
