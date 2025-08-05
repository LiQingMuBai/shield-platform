
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="ID:" prop="id">
    <el-input v-model.number="formData.id" :clearable="true" placeholder="请输入id字段" />
</el-form-item>
        <el-form-item label="username字段:" prop="username">
    <el-input v-model="formData.username" :clearable="true" placeholder="请输入username字段" />
</el-form-item>
        <el-form-item label="associates字段:" prop="associates">
    <el-input v-model="formData.associates" :clearable="true" placeholder="请输入associates字段" />
</el-form-item>
        <el-form-item label="backupChatId字段:" prop="backupChatId">
    <el-input v-model="formData.backupChatId" :clearable="true" placeholder="请输入backupChatId字段" />
</el-form-item>
        <el-form-item label="金额:" prop="amount">
    <el-input v-model="formData.amount" :clearable="true" placeholder="请输入amount字段" />
</el-form-item>
        <el-form-item label="tronAmount字段:" prop="tronAmount">
    <el-input v-model="formData.tronAmount" :clearable="true" placeholder="请输入tronAmount字段" />
</el-form-item>
        <el-form-item label="tronAddress字段:" prop="tronAddress">
    <el-input v-model="formData.tronAddress" :clearable="true" placeholder="请输入tronAddress字段" />
</el-form-item>
        <el-form-item label="ethAddress字段:" prop="ethAddress">
    <el-input v-model="formData.ethAddress" :clearable="true" placeholder="请输入ethAddress字段" />
</el-form-item>
        <el-form-item label="ethAmount字段:" prop="ethAmount">
    <el-input v-model="formData.ethAmount" :clearable="true" placeholder="请输入ethAmount字段" />
</el-form-item>
        <el-form-item label="创建时间:" prop="createdAt">
    <el-date-picker v-model="formData.createdAt" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
        <el-form-item label="deadline字段:" prop="deadline">
    <el-date-picker v-model="formData.deadline" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
        <el-form-item label="修改时间:" prop="updatedAt">
    <el-date-picker v-model="formData.updatedAt" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
        <el-form-item label="userId字段:" prop="userId">
    <el-input v-model="formData.userId" :clearable="true" placeholder="请输入userId字段" />
</el-form-item>
        <el-form-item label="预警次数:" prop="times">
    <el-input v-model.number="formData.times" :clearable="true" placeholder="请输入times字段" />
</el-form-item>
        <el-form-item label="地址:" prop="address">
    <el-input v-model="formData.address" :clearable="true" placeholder="请输入address字段" />
</el-form-item>
        <el-form-item label="privateKey字段:" prop="privateKey">
    <el-input v-model="formData.privateKey" :clearable="true" placeholder="请输入privateKey字段" />
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
  createTgUsers,
  updateTgUsers,
  findTgUsers
} from '@/api/ushield/tgUsers'

defineOptions({
    name: 'TgUsersForm'
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
            username: '',
            associates: '',
            backupChatId: '',
            amount: '',
            tronAmount: '',
            tronAddress: '',
            ethAddress: '',
            ethAmount: '',
            createdAt: new Date(),
            deadline: new Date(),
            updatedAt: new Date(),
            userId: '',
            times: undefined,
            address: '',
            privateKey: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findTgUsers({ ID: route.query.id })
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
               res = await createTgUsers(formData.value)
               break
             case 'update':
               res = await updateTgUsers(formData.value)
               break
             default:
               res = await createTgUsers(formData.value)
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
