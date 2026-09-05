<template>
  <a-modal
    :open="open"
    title="菜单操作权限配置"
    :width="1120"
    :confirm-loading="confirmLoading"
    :mask-closable="false"
    ok-text="保存"
    cancel-text="取消"
    @ok="handleOk"
    @cancel="handleCancel"
  >
    <a-spin :spinning="loading">
      <a-form :label-col="{ span: 3 }" :wrapper-col="{ span: 20 }">
        <a-form-item label="菜单">
          <a-input :value="detailData?.name" disabled style="width: 260px" />
          <div class="permission-tip">
            权限点代表该页面上的一个可授权操作：唯一键与前端按钮权限码一一对应，并绑定该操作需要访问的接口；授权后按钮可见且对应接口可访问。保存后权限点与接口绑定一次性生效
            <br />
            动作类型固定为「查看 / 编辑 / 删除」三类（新增、绑定、重置等写操作统一归到“编辑”），唯一键按类型自动生成，可再手工修改
          </div>
        </a-form-item>
        <a-form-item label="权限动作">
          <div v-if="permissions.length > 0" class="permission-head">
            <span class="col-type">动作类型</span>
            <span class="col-name">权限名称</span>
            <span class="col-key">唯一键</span>
            <span class="col-switch">状态</span>
            <span class="col-api">接口</span>
          </div>
          <div v-for="(item, index) in permissions" :key="index" class="permission-row">
            <a-select
              v-model:value="item.type"
              :options="PERMISSION_TYPE_OPTIONS"
              style="width: 112px"
              @change="onTypeChange(item)"
            />
            <a-input
              v-model:value="item.name"
              placeholder="如 重置密码"
              style="width: 170px"
              @input="markManual(item)"
            />
            <a-input
              v-model:value="item.key"
              placeholder="如 AdminUserResetPwd"
              style="width: 250px"
              @input="markManual(item)"
            />
            <a-switch
              v-model:checked="item.enabled"
              checked-children="启用"
              un-checked-children="禁用"
            />
            <a-button
              size="small"
              :type="item.apiIds?.length ? 'primary' : 'default'"
              @click="openApiBind(index)"
            >
              <template #icon><ApiOutlined /></template>
              {{ item.apiIds?.length ? `已绑 ${item.apiIds.length}` : '绑定接口' }}
            </a-button>
            <a-button type="link" size="small" danger @click="removeRow(index)">删除</a-button>
          </div>
          <div class="permission-actions">
            <a-button type="dashed" block @click="addRow">
              <template #icon><PlusOutlined /></template>
              添加权限动作
            </a-button>
          </div>
          <a-empty v-if="permissions.length === 0" description="暂无权限配置，可点击上方按钮添加" />
        </a-form-item>
      </a-form>
    </a-spin>
    <PermissionApiBindModal
      v-model:open="apiBindOpen"
      :api-ids="currentApiIds"
      :permission-name="currentPermissionName"
      @ok="onApiBindOk"
    />
  </a-modal>
</template>

<script setup lang="ts">
import { computed, ref, watch } from 'vue';
import { message } from 'ant-design-vue';
import { ApiOutlined, PlusOutlined } from '@ant-design/icons-vue';
import { getAdminMenuPermissions } from '@/api/admin/menu';
import { addAdminMenuPermissions } from '@/api/admin/permission';
import PermissionApiBindModal from './PermissionApiBindModal.vue';
import {
  DEFAULT_PERMISSION_TYPES,
  DEFAULT_PERMISSION_TEMPLATE_TYPES,
  PERMISSION_TYPE_OPTIONS,
  handleKey,
} from '@/views/admin/permission/components/common';
import type { MenuTreeItem, MenuPermissionItem } from '@/types/admin_menu';

const props = defineProps<{
  open: boolean;
  detailData?: MenuTreeItem;
}>();
const emit = defineEmits<{ (e: 'update:open', value: boolean): void }>();

const loading = ref(false);
const confirmLoading = ref(false);
/** 权限点行；auto 标记行由系统自动生成 key/name，未人工改动时随类型联动 */
type PermissionRow = MenuPermissionItem & { auto?: boolean };
const permissions = ref<PermissionRow[]>([]);
/** 接口绑定弹窗：当前编辑的权限点行下标 */
const apiBindOpen = ref(false);
const apiBindIndex = ref(-1);
const currentRow = computed(() =>
  apiBindIndex.value >= 0 ? permissions.value[apiBindIndex.value] : undefined,
);
const currentApiIds = computed(() => currentRow.value?.apiIds || []);
const currentPermissionName = computed(() => currentRow.value?.name || '');

const keyPattern = new RegExp('^([A-Z][a-zA-Z0-9]*)+$');

function typeNameOf(type: string) {
  return DEFAULT_PERMISSION_TYPES.find((t) => t.key === type)?.name || type;
}

watch(
  () => props.open,
  (val) => {
    if (val && props.detailData) {
      permissions.value = [];
      loading.value = true;
      // 先取出菜单信息，避免异步回调中 props 变化导致取值异常
      const menuId = props.detailData.id;
      const menuName = props.detailData.name || '';
      const menuPath = props.detailData.path || '';
      getAdminMenuPermissions({ menuId })
        .then((res) => {
          const existing = res.data.permissions || [];
          if (existing.length > 0) {
            // 已有配置：名称/唯一键交由人工维护，不再随类型联动；接口绑定随行带回用于回显
            permissions.value = existing.map((item) => ({
              ...item,
              apiIds: item.apiIds || [],
              auto: false,
            }));
            return;
          }
          // 菜单尚未配置权限时，按默认模板生成查看/新增/编辑/删除四类基础动作
          permissions.value = DEFAULT_PERMISSION_TEMPLATE_TYPES.map((type) => ({
            menuId,
            type: type.key,
            name: `${menuName}${type.name}`,
            key: handleKey(menuPath, type.key),
            enabled: true,
            describe: '',
            apiIds: [],
            auto: true,
          } as PermissionRow));
        })
        .finally(() => {
          loading.value = false;
        });
    }
  },
);

function addRow() {
  permissions.value.push({
    menuId: props.detailData?.id || 0,
    type: 'view',
    name: `${props.detailData?.name || ''}${typeNameOf('view')}`,
    key: handleKey(props.detailData?.path || '', 'view'),
    enabled: true,
    describe: '',
    apiIds: [],
    auto: true,
  });
}

/** 打开接口绑定弹窗（勾选结果暂存到行上，随"保存"一起提交） */
function openApiBind(index: number) {
  apiBindIndex.value = index;
  apiBindOpen.value = true;
}

function onApiBindOk(apiIds: number[]) {
  const row = currentRow.value;
  if (row) {
    row.apiIds = apiIds;
  }
  apiBindOpen.value = false;
}

function onTypeChange(item: PermissionRow) {
  // 自动生成的行跟随类型联动 key/name；人工维护过的行只更新类型
  const type = item.type || 'view';
  if (item.auto) {
    item.key = handleKey(props.detailData?.path || '', type);
    item.name = `${props.detailData?.name || ''}${typeNameOf(type)}`;
  }
}

function markManual(item: PermissionRow) {
  item.auto = false;
}

function removeRow(index: number) {
  permissions.value.splice(index, 1);
}

function handleOk() {
  if (!props.detailData) return;
  for (const item of permissions.value) {
    const name = (item.name || '').trim();
    const key = (item.key || '').trim();
    if (!name || !key) {
      message.warning('请完整填写权限名称与唯一键');
      return;
    }
    if (!keyPattern.test(key)) {
      message.warning(`唯一键 ${key} 请按照驼峰法命名`);
      return;
    }
    item.name = name;
    item.key = key;
  }
  confirmLoading.value = true;
  const list = permissions.value.map((item) => ({
    menuId: props.detailData?.id,
    menuName: props.detailData?.name,
    menuPath: props.detailData?.path,
    id: item.id,
    name: item.name,
    key: item.key,
    type: item.type,
    describe: item.describe,
    enabled: item.enabled,
    apiIds: item.apiIds || [],
  }));
  addAdminMenuPermissions(list)
    .then((res) => {
      message.success(res.msg, 2);
      emit('update:open', false);
    })
    .finally(() => {
      confirmLoading.value = false;
    });
}

function handleCancel() {
  emit('update:open', false);
}
</script>

<style scoped lang="less">
.permission-head,
.permission-row {
  display: flex;
  align-items: center;
  gap: 12px;
}

.permission-head {
  margin-bottom: 8px;
  color: rgba(0, 0, 0, 0.45);
  font-size: 12px;

  .col-type {
    flex: 0 0 112px;
  }

  .col-name {
    flex: 0 0 170px;
  }

  .col-key {
    flex: 1 1 auto;
    min-width: 250px;
  }

  .col-switch {
    flex: 0 0 60px;
  }

  .col-api {
    flex: 0 0 140px;
  }
}

.permission-row {
  margin-bottom: 8px;

  .ant-input {
    flex: none;
  }
}

.permission-actions {
  margin-top: 4px;
}

.permission-tip {
  margin-top: 4px;
  color: rgba(0, 0, 0, 0.45);
  font-size: 12px;
  line-height: 1.5;
}
</style>
