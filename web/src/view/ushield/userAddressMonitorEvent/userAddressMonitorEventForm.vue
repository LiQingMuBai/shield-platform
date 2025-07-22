
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="id字段:" prop="id">
    <el-input v-model.number="formData.id" :clearable="true" placeholder="请输入id字段" />
</el-form-item>
        <el-form-item label="创建时间:" prop="createdAt">
    <el-date-picker v-model="formData.createdAt" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
        <el-form-item label="updatedAt字段:" prop="updatedAt">
    <el-date-picker v-model="formData.updatedAt" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
        <el-form-item label="deletedAt字段:" prop="deletedAt">
    <el-date-picker v-model="formData.deletedAt" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
        <el-form-item label="days字段:" prop="days">
    <el-input v-model.number="formData.days" :clearable="true" placeholder="请输入days字段" />
</el-form-item>
        <el-form-item label="chatId字段:" prop="chatId">
    <el-input v-model.number="formData.chatId" :clearable="true" placeholder="请输入chatId字段" />
</el-form-item>
        <el-form-item label="status字段:" prop="status">
    <el-input v-model.number="formData.status" :clearable="true" placeholder="请输入status字段" />
</el-form-item>
        <el-form-item label="network字段:" prop="network">
    <el-input v-model="formData.network" :clearable="true" placeholder="请输入network字段" />
</el-form-item>
        <el-form-item label="address字段:" prop="address">
    <el-input v-model="formData.address" :clearable="true" placeholder="请输入address字段" />
</el-form-item>
        <el-form-item>
          <el-button :loading="btnLoading" type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createUserAddressMonitorEvent,
  updateUserAddressMonitorEvent,
  findUserAddressMonitorEvent
} from '@/api/ushield/userAddressMonitorEvent'

defineOptions({
    name: 'UserAddressMonitorEventForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'


const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

const type = ref('')
const formData = ref({
            id: undefined,
            createdAt: new Date(),
            updatedAt: new Date(),
            deletedAt: new Date(),
            days: undefined,
            chatId: undefined,
            status: undefined,
            network: '',
            address: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findUserAddressMonitorEvent({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      btnLoading.value = true
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return btnLoading.value = false
            let res
           switch (type.value) {
             case 'create':
               res = await createUserAddressMonitorEvent(formData.value)
               break
             case 'update':
               res = await updateUserAddressMonitorEvent(formData.value)
               break
             default:
               res = await createUserAddressMonitorEvent(formData.value)
               break
           }
           btnLoading.value = false
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
