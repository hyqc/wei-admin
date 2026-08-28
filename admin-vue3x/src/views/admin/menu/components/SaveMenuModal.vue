<template>
  <a-modal
    :open="open"
    :title="detailData?.id ? '编辑菜单' : '新建菜单'"
    :width="DefaultModalWidth"
    :confirm-loading="confirmLoading"
    :mask-closable="false"
    ok-text="保存"
    cancel-text="取消"
    @ok="handleOk"
    @cancel="handleCancel"
  >
    <a-form ref="formRef" :model="formState" :rules="rules" :label-col="{ span: 8 }" :wrapper-col="{ span: 14 }">
      <a-form-item label="父级菜单" name="parentId">
        <a-tree-select
          v-model:value="formState.parentId"
          :tree-data="treeSelectData"
          tree-default-expand-all
          placeholder="请选择父级菜单"
          allow-clear
        />
      </a-form-item>
      <a-form-item label="菜单名称" name="name">
        <a-input v-model:value="formState.name" placeholder="请输入菜单名称" allow-clear />
      </a-form-item>
      <a-form-item label="菜单路径" name="path">
        <a-input v-model:value="formState.path" placeholder="请输入菜单路径" allow-clear @change="onPathChange" />
      </a-form-item>
      <a-form-item label="键名" name="key">
        <a-input v-model:value="formState.key" disabled />
      </a-form-item>
      <a-form-item label="排序" name="sort">
        <a-input-number v-model:value="formState.sort" :min="0" style="width: 100%" />
      </a-form-item>
      <a-form-item label="图标" name="icon">
        <a-input v-model:value="formState.icon" placeholder="请输入图标名称" allow-clear />
      </a-form-item>
      <a-form-item label="重定向地址" name="redirect">
        <a-input v-model:value="formState.redirect" placeholder="请输入重定向地址" allow-clear />
      </a-form-item>
      <a-form-item label="描述" name="describe">
        <a-input v-model:value="formState.describe" placeholder="请输入菜单描述" allow-clear />
      </a-form-item>
      <a-form-item label="是否隐藏菜单" name="isHideInMenu">
        <a-switch v-model:checked="formState.isHideInMenu" checked-children="隐藏" un-checked-children="显示" />
      </a-form-item>
      <a-form-item label="是否隐藏子菜单" name="isHideChildrenInMenu">
        <a-switch v-model:checked="formState.isHideChildrenInMenu" checked-children="隐藏" un-checked-children="显示" />
      </a-form-item>
      <a-form-item label="是否启用" name="isEnabled">
        <a-switch v-model:checked="formState.isEnabled" checked-children="启用" un-checked-children="禁用" />
      </a-form-item>
    </a-form>
  </a-modal>
</template>

<script setup lang="ts">
import { computed, reactive, ref, watch } from 'vue';
import { message } from 'ant-design-vue';
import { addAdminMenu, editAdminMenu } from '@/api/admin/menu';
import { AdminMenuKey, AdminRouterPath } from '@/api/pattern';
import { DefaultModalWidth } from '@/api/config';
import type { MenuTreeItem, ReqAdminMenuAdd } from '@/types/admin_menu';
import type { FormInstance } from 'ant-design-vue';

const props = defineProps<{
  open: boolean;
  tree: MenuTreeItem[];
  detailData?: MenuTreeItem;
}>();
const emit = defineEmits<{ (e: 'update:open', value: boolean): void; (e: 'notice'): void }>();

const formRef = ref<FormInstance>();
const confirmLoading = ref(false);

const formState = reactive<{
  parentId?: number;
  name: string;
  path: string;
  key: string;
  sort?: number;
  icon?: string;
  redirect?: string;
  describe?: string;
  isHideInMenu: boolean;
  isHideChildrenInMenu: boolean;
  isEnabled: boolean;
}>({
  parentId: undefined,
  name: '',
  path: '',
  key: '',
  sort: 0,
  icon: '',
  redirect: '/',
  describe: '',
  isHideInMenu: false,
  isHideChildrenInMenu: false,
  isEnabled: true,
});

const rules = {
  name: [
    { required: true, message: '请输入菜单名称' },
    { max: 50, message: '菜单名称长度不能超过50个字符' },
  ],
  path: [
    { required: true, message: '请输入菜单路径' },
    { pattern: AdminRouterPath, message: '路径格式不正确' },
  ],
  key: [
    { required: true, pattern: AdminMenuKey, message: '请按照驼峰法命名' },
  ],
};

/** 把路径转为键名 */
function path2UpperCamelCase(path: string) {
  return path
    ?.split('/')
    .filter((name) => name.length > 0)
    .map((name) => name[0].toUpperCase() + name.substring(1))
    .join('');
}

/** 构建树选择数据 */
const treeSelectData = computed(() => {
  const build = (list: MenuTreeItem[]): any[] =>
    list.map((item) => ({
      key: item.id,
      value: item.id,
      title: item.name,
      children: item.children?.length ? build(item.children) : undefined,
    }));
  return build(props.tree);
});

function onPathChange() {
  formState.key = path2UpperCamelCase(formState.path);
}

watch(
  () => props.open,
  (val) => {
    if (val) {
      if (props.detailData?.id) {
        // 编辑
        formState.parentId = props.detailData.parentId || 0;
        formState.name = props.detailData.name || '';
        formState.path = props.detailData.path || '';
        formState.key = props.detailData.key || '';
        formState.sort = props.detailData.sort ?? 0;
        formState.icon = props.detailData.icon || '';
        formState.redirect = props.detailData.redirect || '/';
        formState.describe = props.detailData.describe || '';
        formState.isHideInMenu = !!props.detailData.hideInMenu;
        formState.isHideChildrenInMenu = !!props.detailData.hideChildrenInMenu;
        formState.isEnabled = !!props.detailData.enabled;
      } else {
        // 新增：指定父级
        formRef.value?.resetFields();
        formState.parentId = props.detailData?.id;
        formState.isEnabled = true;
        formState.isHideInMenu = false;
        formState.isHideChildrenInMenu = false;
        formState.redirect = '/';
      }
    }
  },
);

function handleOk() {
  formRef.value
    ?.validate()
    .then(() => {
      const data: ReqAdminMenuAdd & { id?: number } = {
        ...formState,
        parentId: formState.parentId ?? 0,
      };
      if (props.detailData?.id) {
        // 编辑
        data.id = props.detailData.id;
        confirmLoading.value = true;
        editAdminMenu(data)
          .then((res) => {
            message.success(res.msg, 2);
            emit('notice');
            emit('update:open', false);
          })
          .finally(() => {
            confirmLoading.value = false;
          });
      } else {
        confirmLoading.value = true;
        addAdminMenu(data)
          .then((res) => {
            message.success(res.msg, 2);
            emit('notice');
            emit('update:open', false);
          })
          .finally(() => {
            confirmLoading.value = false;
          });
      }
    })
    .catch(() => {});
}

function handleCancel() {
  emit('update:open', false);
}
</script>
