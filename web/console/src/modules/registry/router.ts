import { Router } from '../../../helpers/Router';

/**
 * @param sub   二级菜单栏的一级导航
 * @param mode  当前的展示内容类型 list | create | update | detail
 * @param tab   详情页 tab
 * @param name  镜像仓库键值
 */
export const router = new Router('/tkestack/registry/appcenter(/:sub)(/:mode)(/:tab)', {
  sub: '',
  mode: '',
  tab: '',
  cg: '',
  chart: '',
  cgName: '',
  chartName: '',
  prj: '',
  ns: '',
});
