<template>

    <div class="app-container">
        <el-row style="margin: 20px">
            <el-col :span="24" style="margin: 20px">
                <div class="grid-content bg-purple-dark">菜单管理</div>
            </el-col>

            <el-col :span="24" style="margin: 20px 0 ">
                <el-button
                        @click="handleCreate"
                        class="filter-item"
                        icon="el-icon-edit"
                        style="margin-left: 10px;"
                        type="primary"
                >新增
                </el-button>
            </el-col>
            <el-col :span="24">
                <el-table
                        :data="tableData"
                        style="width: 100%">
                    <el-table-column
                            label="日期"
                            prop="date"
                            width="180">
                    </el-table-column>
                    <el-table-column
                            label="姓名"
                            prop="name"
                            width="180">
                    </el-table-column>
                    <el-table-column
                            label="地址"
                            prop="address">
                    </el-table-column>
                    <el-table-column align="center" class-name="small-padding fixed-width" label="Actions" width="230">
                        <template slot-scope="{row,$index}">
                            <el-button @click="handleUpdate(row)" size="mini" type="primary">
                                Edit
                            </el-button>
                        </template>
                    </el-table-column>
                </el-table>
            </el-col>
        </el-row>

        <el-dialog
                :visible.sync="dialogFormVisible"
                center
                title="新增菜单"
                width="30%">
            <el-form :model="form" label-width="80px">
                <el-form-item label="菜单名称">
                    <el-input placeholder="请输入菜单名称" v-model="form.name"></el-input>
                </el-form-item>
                <el-form-item label="菜单路径">
                    <el-input clearable placeholder="例：/api/home" v-model="form.name"></el-input>
                </el-form-item>
                <el-form-item label="是否启用">
                    <el-switch v-model="form.name"></el-switch>
                </el-form-item>
            </el-form>
            <div style="text-align: center">
                <el-button @click="dialogFormVisible = false">取 消</el-button>
                <el-button @click="dialogFormVisible = false" type="primary">确 定</el-button>
            </div>

        </el-dialog>
    </div>


</template>
<script>
    export default {
        data() {
            return {
                dialogFormVisible: false,
                form: {
                    name: '',
                },
                tableData: [{
                    date: '2016-05-02',
                    name: '王小虎',
                    address: '上海市普陀区金沙江路 1518 弄'
                }, {
                    date: '2016-05-04',
                    name: '王小虎',
                    address: '上海市普陀区金沙江路 1517 弄'
                }, {
                    date: '2016-05-01',
                    name: '王小虎',
                    address: '上海市普陀区金沙江路 1519 弄'
                }, {
                    date: '2016-05-03',
                    name: '王小虎',
                    address: '上海市普陀区金沙江路 1516 弄'
                }]
            };
        },
        methods: {
            handleCreate() {
                this.dialogFormVisible = true;
            }
            // search() {
            //     this.page.pi = 1;
            //     this.init();
            // },

        },
        mounted() {
            // this.init_();
            // this.init();
        }
    };
</script>
