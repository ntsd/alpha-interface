import { Demo, AlphaInterface, Layout } from './pages';

const routes = [
  {
    name: '/',
    component: AlphaInterface,
    layout: Layout,
  },
  {
    name: 'demo',
    component: Demo,
    layout: Layout,
  },
];

export { routes };
