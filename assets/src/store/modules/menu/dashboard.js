import lazyLoading from './lazyLoading';
import {ROUTE_DASHBOARD} from './../../../router/path';

export default {
  name: 'Dashboard',
  path: ROUTE_DASHBOARD,
  component: lazyLoading('dashboard/Dashboard'),
  meta: {
    default: true,
    auth: true,
    title: 'Dashboard',
    iconClass: 'vuestic-icon vuestic-icon-dashboard'
  }
};
