<template>
  <PageContainer title="菜单管理">
    <template #extra>
      <a-button type="primary" @click="openAddModal()">
        <template #icon><PlusOutlined /></template>
        新建菜单
      </a-button>
    </template>

    <a-table
      :columns="columns"
      :data-source="rows"
      :loading="loading"
      :pagination="false"
      :expanded-row-keys="expandedRowKeys"
      @update:expanded-row-keys="(keys: number[]) => (expandedRowKeys = keys)"
      row-key="id"
      :scroll="{ x: 1160 }"
    >
      <template #bodyCell="{ column, record }">
        <template v-if="column.key === 'name'">
          <a-space>
            <span>{{ record.name }}</span>
          </a-space>
        </template>
        <template v-else-if="column.key === 'isEnabled'">
          <a-popconfirm
            :title="`确定要${record.enabled ? '禁用' : '启用'}该菜单吗？`"
            ok-text="确定"
            cancel-text="取消"
            @confirm="updateEnabled(record)"
          >
            <a-switch :checked="record.enabled" checked-children="启用" un-checked-children="禁用" />
          </a-popconfirm>
        </template>
        <template v-else-if="column.key === 'hideInMenu'">
          <a-switch
            :checked="!record.hideInMenu"
            checked-children="显示"
            un-checked-children="隐藏"
            @change="updateShow(record)"
          />
        </template>
        <template v-else-if="column.key === 'action'">
          <a-space>
            <Authorization permission="AdminMenuEdit">
              <a-button type="link" size="small" @click="openAddModal(record)">
                <template #icon><PlusOutlined /></template>
                新增子菜单
              </a-button>
            </Authorization>
            <Authorization permission="AdminMenuEdit">
              <a-button type="link" size="small" @click="openEditModal(record)">
                <template #icon><EditOutlined /></template>
                编辑
              </a-button>
            </Authorization>
            <Authorization permission="AdminMenuView">
              <a-button type="link" size="small" @click="openDetailModal(record)">
                <template #icon><EyeOutlined /></template>
                详情
              </a-button>
            </Authorization>
            <Authorization permission="AdminMenuEdit">
              <a-button type="link" size="small" @click="openPermissionsModal(record)">
                <template #icon><SafetyCertificateOutlined /></template>
                权限配置
              </a-button>
            </Authorization>
            <Authorization permission="AdminMenuDelete">
              <a-popconfirm
                v-if="!record.enabled"
                title="确定要删除该菜单吗？"
                ok-text="确定"
                cancel-text="取消"
                @confirm="onDelete(record)"
              >
                <a-button type="link" size="small" danger>
                  <template #icon><DeleteOutlined /></template>
                  删除
                </a-button>
              </a-popconfirm>
            </Authorization>
          </a-space>
        </template>
      </template>
    </a-table>

    <SaveMenuModal
      v-model:open="saveModalStatus"
      :tree="rows"
      :detail-data="saveData"
      @notice="onModalNotice"
    />
    <DetailMenuDrawer v-model:open="detailModalStatus" :detail-data="detailData" />
    <PermissionsSaveModal v-model:open="permissionsModalStatus" :detail-data="detailData" />
  </PageContainer>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue';
import { message } from 'ant-design-vue';
import {
  PlusOutlined,
  EyeOutlined,
  EditOutlined,
  SafetyCertificateOutlined,
  DeleteOutlined,
} from '@ant-design/icons-vue';
import PageContainer from '@/components/PageContainer.vue';
import Authorization from '@/components/Authorization.vue';
import SaveMenuModal from './components/SaveMenuModal.vue';
import DetailMenuDrawer from './components/DetailMenuDrawer.vue';
import PermissionsSaveModal from './components/PermissionsSaveModal.vue';
import { getAdminMenuTree, deleteAdminMenu, enableAdminMenu, showAdminMenu } from '@/api/admin/menu';
import type { MenuTreeItem } from '@/types/admin_menu';

const loading = ref(false);
const rows = ref<MenuTreeItem[]>([]);
/** 树形表格展开的行 key，加载后默认全部展开 */
const expandedRowKeys = ref<number[]>([]);

const saveModalStatus = ref(false);
const detailModalStatus = ref(false);
const permissionsModalStatus = ref(false);
const saveData = ref<MenuTreeItem>();
const detailData = ref<MenuTreeItem>();

const columns = [
  { title: '菜单名称', dataIndex: 'name', key: 'name', width: 220 },
  { title: '键名', dataIndex: 'key', key: 'key', width: 160 },
  { title: '菜单路径', dataIndex: 'path', key: 'path', width: 180 },
  { title: '排序', dataIndex: 'sort', key: 'sort', width: 80 },
  { title: '状态', dataIndex: 'enabled', key: 'isEnabled', width: 100 },
  { title: '是否显示', dataIndex: 'hideInMenu', key: 'hideInMenu', width: 100 },
  { title: '创建时间', dataIndex: 'createTime', key: 'createTime', width: 180, customRender: ({ text }: { text: number }) => formatTime(text) },
  { title: '操作', dataIndex: 'action', key: 'action', fixed: 'right', width: 380 },
];

function formatTime(time?: number) {
  if (!time) return '';
  const d = new Date(time * 1000);
  const pad = (n: number) => String(n).padStart(2, '0');
  return `${d.getFullYear()}-${pad(d.getMonth() + 1)}-${pad(d.getDate())} ${pad(d.getHours())}:${pad(d.getMinutes())}:${pad(d.getSeconds())}`;
}

/** 递归收集所有含子节点的菜单 id */
function collectParentKeys(list: MenuTreeItem[]): number[] {
  const keys: number[] = [];
  for (const item of list) {
    if (item.children?.length) {
      keys.push(item.id as number);
      keys.push(...collectParentKeys(item.children));
    }
  }
  return keys;
}

function getRows() {
  loading.value = true;
  getAdminMenuTree()
    .then((res) => {
      rows.value = res.data.list || [];
      expandedRowKeys.value = collectParentKeys(rows.value);
    })
    .finally(() => {
      loading.value = false;
    });
}

function updateEnabled(record: MenuTreeItem) {
  enableAdminMenu({ menuId: record.id, enabled: !record.enabled }).then((res) => {
    message.success(res.msg, 2);
    getRows();
  });
}

function updateShow(record: MenuTreeItem) {
  showAdminMenu({ menuId: record.id, show: !!record.hideInMenu }).then((res) => {
    message.success(res.msg, 2);
    getRows();
  });
}

function onDelete(record: MenuTreeItem) {
  deleteAdminMenu({ menuId: record.id }).then((res) => {
    message.success(res.msg, 2);
    getRows();
  });
}

function openAddModal(parent?: MenuTreeItem) {
  saveData.value = parent ? { ...parent } : undefined;
  saveModalStatus.value = true;
}

function openEditModal(record: MenuTreeItem) {
  saveData.value = { ...record };
  saveModalStatus.value = true;
}

function openDetailModal(record: MenuTreeItem) {
  detailData.value = record;
  detailModalStatus.value = true;
}

function openPermissionsModal(record: MenuTreeItem) {
  detailData.value = record;
  permissionsModalStatus.value = true;
}

function onModalNotice() {
  getRows();
}

onMounted(() => {
  getRows();
});
</script>
