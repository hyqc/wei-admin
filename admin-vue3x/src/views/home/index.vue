<template>
  <div class="home-page">
    <a-card title="排班日历">
      <a-calendar v-model:value="currentDate" fullscreen>
        <template #dateCellRender="{ current }">
          <ul class="events">
            <li v-if="current.date() === today.date()" class="event-today">
              <a-tag color="blue">今日无排班计划</a-tag>
            </li>
          </ul>
        </template>
      </a-calendar>
    </a-card>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import dayjs, { type Dayjs } from 'dayjs';

const today = dayjs();
const currentDate = ref<Dayjs>(dayjs());
</script>

<style scoped lang="less">
.home-page {
  .events {
    margin: 0;
    padding: 0;
    list-style: none;

    .event-today {
      margin-top: 4px;
    }
  }
}
</style>
