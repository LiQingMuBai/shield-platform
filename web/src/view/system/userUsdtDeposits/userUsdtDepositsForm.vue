
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
        <el-form-item label="userId字段:" prop="userId">
    <el-input v-model.number="formData.userId" :clearable="true" placeholder="请输入userId字段" />
</el-form-item>
        <el-form-item label="status字段:" prop="status">
    <el-input v-model.number="formData.status" :clearable="true" placeholder="请输入status字段" />
</el-form-item>
        <el-form-item label="placeholder字段:" prop="placeholder">
    <el-input v-model="formData.placeholder" :clearable="true" placeholder="请输入placeholder字段" />
</el-form-item>
        <el-form-item label="地址:" prop="address">
    <el-input v-model="formData.address" :clearable="true" placeholder="请输入地址" />
</el-form-item>
        <el-form-item label="tx_hash:" prop="txHash">
    <el-input v-model="formData.txHash" :clearable="true" placeholder="请输入tx_hash" />
</el-form-item>
        <el-form-item label="金额:" prop="amount">
    <el-input v-model="formData.amount" :clearable="true" placeholder="请输入金额" />
</el-form-item>
        <el-form-item label="区块:" prop="block">
    <el-input v-model="formData.block" :clearable="true" placeholder="请输入区块" />
</el-form-item>
        <el-form-item label="orderNo字段:" prop="orderNo">
    <el-input v-model="formData.orderNo" :clearable="true" placeholder="请输入orderNo字段" />
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
  createUserUsdtDeposits,
  updateUserUsdtDeposits,
  findUserUsdtDeposits
} from '@/api/system/userUsdtDeposits'

defineOptions({
    name: 'UserUsdtDepositsForm'
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
            userId: undefined,
            status: undefined,
            placeholder: '',
            address: '',
            txHash: '',
            amount: '',
            block: '',
            orderNo: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findUserUsdtDeposits({ ID: route.query.id })
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
               res = await createUserUsdtDeposits(formData.value)
               break
             case 'update':
               res = await updateUserUsdtDeposits(formData.value)
               break
             default:
               res = await createUserUsdtDeposits(formData.value)
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
