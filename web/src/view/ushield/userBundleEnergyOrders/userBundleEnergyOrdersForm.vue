
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="ID:" prop="id">
    <el-input v-model.number="formData.id" :clearable="true" placeholder="请输入id字段" />
</el-form-item>
        <el-form-item label="订单ID:" prop="orderNo">
    <el-input v-model="formData.orderNo" :clearable="true" placeholder="请输入orderNo字段" />
</el-form-item>
        <el-form-item label="交易ID:" prop="txId">
    <el-input v-model="formData.txId" :clearable="true" placeholder="请输入txId字段" />
</el-form-item>
        <el-form-item label="from地址:" prop="fromAddress">
    <el-input v-model="formData.fromAddress" :clearable="true" placeholder="请输入fromAddress字段" />
</el-form-item>
        <el-form-item label="to地址:" prop="toAddress">
    <el-input v-model="formData.toAddress" :clearable="true" placeholder="请输入toAddress字段" />
</el-form-item>
        <el-form-item label="代币:" prop="token">
    <el-input v-model="formData.token" :clearable="true" placeholder="请输入token字段" />
</el-form-item>
        <el-form-item label="金额:" prop="amount">
    <el-input-number v-model="formData.amount" style="width:100%" :precision="2" :clearable="true" />
</el-form-item>
        <el-form-item label="飞机ID:" prop="chatId">
    <el-input v-model="formData.chatId" :clearable="true" placeholder="请输入chatId字段" />
</el-form-item>
        <el-form-item label="备注:" prop="remark">
    <el-input v-model="formData.remark" :clearable="true" placeholder="请输入remark字段" />
</el-form-item>
        <el-form-item label="创建时间:" prop="createdAt">
    <el-date-picker v-model="formData.createdAt" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>

        <el-form-item label="修改时间:" prop="updatedAt">
    <el-date-picker v-model="formData.updatedAt" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
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
  createUserBundleEnergyOrders,
  updateUserBundleEnergyOrders,
  findUserBundleEnergyOrders
} from '@/api/ushield/userBundleEnergyOrders'

defineOptions({
    name: 'UserBundleEnergyOrdersForm'
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
            orderNo: '',
            txId: '',
            fromAddress: '',
            toAddress: '',
            token: '',
            amount: 0,
            chatId: '',
            remark: '',
            createdAt: new Date(),
            deletedAt: new Date(),
            updatedAt: new Date(),
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findUserBundleEnergyOrders({ ID: route.query.id })
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
               res = await createUserBundleEnergyOrders(formData.value)
               break
             case 'update':
               res = await updateUserBundleEnergyOrders(formData.value)
               break
             default:
               res = await createUserBundleEnergyOrders(formData.value)
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
