<template>
  <PageContainer
    :page-info="pageInfo"
    @page-change="onPageChange"
    @page-size-change="onPageSizeChange"
  >
    <template #searchArea>
      <a-form layout="inline" :model="searchForm" class="search-form" @finish="onSearch">
        <a-form-item label="分组" name="uploadGroup">
          <a-input v-model:value="searchForm.uploadGroup" placeholder="如 /admin/user" allow-clear />
        </a-form-item>
        <a-form-item label="文件名" name="originalName">
          <a-input v-model:value="searchForm.originalName" placeholder="原始文件名" allow-clear />
        </a-form-item>
        <a-form-item label="类型" name="ext">
          <a-input v-model:value="searchForm.ext" placeholder="如 png" allow-clear style="width: 100px" />
        </a-form-item>
        <a-form-item label="存储" name="driver">
          <a-select
            v-model:value="searchForm.driver"
            placeholder="全部"
            allow-clear
            :options="UPLOAD_DRIVER_OPTIONS"
            style="width: 140px"
          />
        </a-form-item>
        <a-form-item>
          <a-space>
            <a-button type="primary" html-type="submit">
              <template #icon><SearchOutlined /></template>
              查询
            </a-button>
            <a-button @click="onReset">
              <template #icon><ReloadOutlined /></template>
              重置
            </a-button>
          </a-space>
        </a-form-item>
      </a-form>
    </template>
    <template #extra>
      <Authorization permission="AdminUploadEdit">
        <a-button type="primary" @click="openUploadModal">
          <template #icon><UploadOutlined /></template>
          上传文件
        </a-button>
      </Authorization>
    </template>

    <a-table
      :columns="columns"
      :data-source="rows"
      :loading="loading"
      :pagination="false"
      row-key="id"
      :scroll="{ x: 1400 }"
    >
      <template #bodyCell="{ column, record }">
        <template v-if="column.key === 'file'">
          <a-space>
            <a-image
              v-if="isImage(record)"
              :src="record.url"
              :width="36"
              :height="36"
              style="object-fit: cover; border-radius: 4px"
              :preview="{ mask: false }"
              :fallback="record.url"
            />
            <div class="file-cell">
              <div class="file-name" :title="record.originalName">{{ record.originalName }}</div>
              <div class="file-sub">{{ record.newName }}</div>
            </div>
          </a-space>
        </template>
        <template v-else-if="column.key === 'uploadGroup'">
          <a-tag color="blue">/{{ record.uploadGroup }}/</a-tag>
        </template>
        <template v-else-if="column.key === 'driver'">
          <a-tag>{{ record.driverText || record.driver }}</a-tag>
        </template>
        <template v-else-if="column.key === 'url'">
          <a :href="record.url" target="_blank" rel="noopener noreferrer">打开</a>
        </template>
        <template v-else-if="column.key === 'action'">
          <a-space>
            <a-button type="link" size="small" @click="copyUrl(record)">
              <template #icon><LinkOutlined /></template>
              复制链接
            </a-button>
            <Authorization permission="AdminUploadDelete">
              <a-popconfirm
                title="删除后存储上的文件也会一并删除，确定要删除吗？"
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

    <UploadFileModal v-model:open="uploadModalStatus" @notice="onModalNotice" />
  </PageContainer>
</template>

<script setup lang="ts">
import { onMounted, reactive, ref } from 'vue';
import { message } from 'ant-design-vue';
import {
  SearchOutlined,
  ReloadOutlined,
  UploadOutlined,
  LinkOutlined,
  DeleteOutlined,
} from '@ant-design/icons-vue';
import PageContainer from '@/components/PageContainer.vue';
import Authorization from '@/components/Authorization.vue';
import UploadFileModal from './components/UploadFileModal.vue';
import { getAdminUploadList, deleteAdminUpload } from '@/api/admin/upload';
import { UPLOAD_DRIVER_OPTIONS } from '@/types/admin_upload';
import type { UploadItem } from '@/types/admin_upload';
import { DEFAULT_PAGE_INFO } from '@/api/config';
import type { PageInfoType } from '@/api/types';

const IMAGE_EXTS = ['png', 'jpg', 'jpeg', 'gif', 'webp', 'bmp', 'svg'];

const loading = ref(false);
const rows = ref<UploadItem[]>([]);
const pageInfo = reactive<PageInfoType>({ ...DEFAULT_PAGE_INFO });
const uploadModalStatus = ref(false);

const searchForm = reactive<{
  uploadGroup?: string;
  originalName?: string;
  ext?: string;
  driver?: string;
}>({
  uploadGroup: undefined,
  originalName: undefined,
  ext: undefined,
  driver: undefined,
});

const columns = [
  { title: 'ID', dataIndex: 'id', key: 'id', fixed: 'left', width: 80 },
  { title: '文件', dataIndex: 'file', key: 'file', width: 280 },
  { title: '分组', dataIndex: 'uploadGroup', key: 'uploadGroup', width: 180 },
  { title: '大小', dataIndex: 'sizeText', key: 'sizeText', width: 100 },
  { title: '存储', dataIndex: 'driver', key: 'driver', width: 120 },
  { title: '访问地址', dataIndex: 'url', key: 'url', width: 90 },
  { title: '上传者', dataIndex: 'adminName', key: 'adminName', width: 120 },
  { title: '上传日期', dataIndex: 'uploadDate', key: 'uploadDate', width: 120 },
  { title: '操作', dataIndex: 'action', key: 'action', fixed: 'right', width: 180 },
];

function isImage(record: UploadItem) {
  return !!record.ext && IMAGE_EXTS.includes(record.ext.toLowerCase());
}

function getRows() {
  loading.value = true;
  getAdminUploadList({
    pageNum: pageInfo.pageNum,
    pageSize: pageInfo.pageSize,
    uploadGroup: searchForm.uploadGroup,
    originalName: searchForm.originalName,
    ext: searchForm.ext,
    driver: searchForm.driver,
  })
    .then((res) => {
      rows.value = res.data.list || [];
      pageInfo.total = res.data.pageInfo?.total || rows.value.length;
    })
    .finally(() => {
      loading.value = false;
    });
}

function onSearch() {
  pageInfo.pageNum = 1;
  getRows();
}

function onReset() {
  searchForm.uploadGroup = undefined;
  searchForm.originalName = undefined;
  searchForm.ext = undefined;
  searchForm.driver = undefined;
  onSearch();
}

function onPageChange(pageNum: number) {
  pageInfo.pageNum = pageNum;
  getRows();
}

function onPageSizeChange(pageSize: number) {
  pageInfo.pageSize = pageSize;
  pageInfo.pageNum = 1;
  getRows();
}

function openUploadModal() {
  uploadModalStatus.value = true;
}

function onModalNotice() {
  getRows();
}

async function copyUrl(record: UploadItem) {
  if (!record.url) return;
  try {
    await navigator.clipboard.writeText(record.url);
    message.success('链接已复制', 2);
  } catch {
    message.warning('复制失败，请手动复制');
  }
}

function onDelete(record: UploadItem) {
  deleteAdminUpload({ id: record.id }).then((res) => {
    message.success(res.msg, 2);
    getRows();
  });
}

onMounted(() => {
  getRows();
});
</script>

<style scoped lang="less">
.search-form {
  .ant-form-item {
    margin-bottom: 8px;
  }
}

.file-cell {
  line-height: 1.4;

  .file-name {
    max-width: 200px;
    overflow: hidden;
    white-space: nowrap;
    text-overflow: ellipsis;
  }

  .file-sub {
    color: rgba(0, 0, 0, 0.45);
    font-size: 12px;
  }
}
</style>
