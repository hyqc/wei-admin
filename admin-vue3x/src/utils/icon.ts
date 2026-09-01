import * as AntIcons from '@ant-design/icons-vue';
import type { Component } from 'vue';

/** ant-design 图标库：图标名称 → 组件 */
export const antIcons = AntIcons as unknown as Record<string, Component>;

/**
 * 可选图标名称列表（仅 Outlined 线性风格，按字母排序）
 * 注：图标组件为对象，工具方法（createFromIconfontCN 等）为函数，按后缀过滤即可区分
 */
export const antIconNames: string[] = Object.keys(antIcons)
  .filter((name) => /Outlined$/.test(name))
  .sort();

/** 按名称获取图标组件，名称为空或不存在时返回 undefined */
export function getAntIcon(name?: string): Component | undefined {
  if (!name) return undefined;
  return antIcons[name];
}
