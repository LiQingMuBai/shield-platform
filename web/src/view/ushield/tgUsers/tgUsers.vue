
<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">

        <template v-if="showAllQuery">
          <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
        </template>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
          <el-button link type="primary" icon="arrow-down" @click="showAllQuery=true" v-if="!showAllQuery">展开</el-button>
          <el-button link type="primary" icon="arrow-up" @click="showAllQuery=false" v-else>收起</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button  type="primary" icon="plus" @click="openDialog()">新增</el-button>
            <el-button  icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
            
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="id"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        
            <el-table-column align="left" label="ID" prop="id" width="120" />

            <el-table-column align="left" label="用户名" prop="username" width="120" />

            <el-table-column align="left" label="ChatID" prop="associates" width="120" />

            <el-table-column align="left" label="第二通知人" prop="backupChatId" width="120" />

            <el-table-column align="left" label="USDT余额" prop="amount" width="180" />

            <el-table-column align="left" label="TRX余额" prop="tronAmount" width="180" />


            <el-table-column align="left" label="创建时间" prop="createdAt" width="180">
   <template #default="scope">{{ formatDate(scope.row.createdAt) }}</template>
</el-table-column>
            <el-table-column align="left" label="剩余次数" prop="times" width="120" />

        <el-table-column align="left" label="操作" fixed="right" :min-width="appStore.operateMinWith">
            <template #default="scope">
            <el-button  type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            <el-button  type="primary" link icon="edit" class="table-button" @click="updateTgUsersFunc(scope.row)">编辑</el-button>
            <el-button   type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
       <template #header>
              <div class="flex justify-between items-center">
                <span class="text-lg">{{type==='create'?'新增':'编辑'}}</span>
                <div>
                  <el-button :loading="btnLoading" type="primary" @click="enterDialog">确 定</el-button>
                  <el-button @click="closeDialog">取 消</el-button>
                </div>
              </div>
            </template>

          <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
            <el-form-item label="id字段:" prop="id">
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
            <el-form-item label="amount字段:" prop="amount">
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
            <el-form-item label="updatedAt字段:" prop="updatedAt">
    <el-date-picker v-model="formData.updatedAt" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
            <el-form-item label="userId字段:" prop="userId">
    <el-input v-model="formData.userId" :clearable="true" placeholder="请输入userId字段" />
</el-form-item>
            <el-form-item label="times字段:" prop="times">
    <el-input v-model.number="formData.times" :clearable="true" placeholder="请输入times字段" />
</el-form-item>
            <el-form-item label="address字段:" prop="address">
    <el-input v-model="formData.address" :clearable="true" placeholder="请输入address字段" />
</el-form-item>
            <el-form-item label="privateKey字段:" prop="privateKey">
    <el-input v-model="formData.privateKey" :clearable="true" placeholder="请输入privateKey字段" />
</el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
            <el-descriptions :column="1" border>
                    <el-descriptions-item label="id字段">
    {{ detailForm.id }}
</el-descriptions-item>
                    <el-descriptions-item label="username字段">
    {{ detailForm.username }}
</el-descriptions-item>
                    <el-descriptions-item label="associates字段">
    {{ detailForm.associates }}
</el-descriptions-item>
                    <el-descriptions-item label="backupChatId字段">
    {{ detailForm.backupChatId }}
</el-descriptions-item>
                    <el-descriptions-item label="amount字段">
    {{ detailForm.amount }}
</el-descriptions-item>
                    <el-descriptions-item label="tronAmount字段">
    {{ detailForm.tronAmount }}
</el-descriptions-item>
                    <el-descriptions-item label="tronAddress字段">
    {{ detailForm.tronAddress }}
</el-descriptions-item>
                    <el-descriptions-item label="ethAddress字段">
    {{ detailForm.ethAddress }}
</el-descriptions-item>
                    <el-descriptions-item label="ethAmount字段">
    {{ detailForm.ethAmount }}
</el-descriptions-item>
                    <el-descriptions-item label="创建时间">
    {{ detailForm.createdAt }}
</el-descriptions-item>
                    <el-descriptions-item label="deadline字段">
    {{ detailForm.deadline }}
</el-descriptions-item>
                    <el-descriptions-item label="updatedAt字段">
    {{ detailForm.updatedAt }}
</el-descriptions-item>
                    <el-descriptions-item label="userId字段">
    {{ detailForm.userId }}
</el-descriptions-item>
                    <el-descriptions-item label="times字段">
    {{ detailForm.times }}
</el-descriptions-item>
                    <el-descriptions-item label="address字段">
    {{ detailForm.address }}
</el-descriptions-item>
                    <el-descriptions-item label="privateKey字段">
    {{ detailForm.privateKey }}
</el-descriptions-item>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
  createTgUsers,
  deleteTgUsers,
  deleteTgUsersByIds,
  updateTgUsers,
  findTgUsers,
  getTgUsersList
} from '@/api/ushield/tgUsers'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import { useAppStore } from "@/pinia"




defineOptions({
    name: 'TgUsers'
})

// 提交按钮loading
const btnLoading = ref(false)
const appStore = useAppStore()

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
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
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getTgUsersList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteTgUsersFunc(row)
        })
    }

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
      const ids = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          ids.push(item.id)
        })
      const res = await deleteTgUsersByIds({ ids })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === ids.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
      })
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateTgUsersFunc = async(row) => {
    const res = await findTgUsers({ id: row.id })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteTgUsersFunc = async (row) => {
    const res = await deleteTgUsers({ id: row.id })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
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
        }
}
// 弹窗确定
const enterDialog = async () => {
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
                closeDialog()
                getTableData()
              }
      })
}

const detailForm = ref({})

// 查看详情控制标记
const detailShow = ref(false)


// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}


// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findTgUsers({ id: row.id })
  if (res.code === 0) {
    detailForm.value = res.data
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  detailForm.value = {}
}


</script>

<style>

</style>
