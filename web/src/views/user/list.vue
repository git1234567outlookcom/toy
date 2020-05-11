<template>
    <el-row style="margin: 20px">
        <el-col :span="24" style="margin: 20px">
            <div class="grid-content bg-purple-dark">用户列表</div>
        </el-col>

        <el-col :span="24" style="margin: 20px 0 ">
            <el-button type="success">新增</el-button>
            <el-button type="primary">修改</el-button>
            <el-button type="warning">修改</el-button>
            <el-button type="danger">删除</el-button>
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
            </el-table>

        </el-col>
    </el-row>
</template>
<script>
    export default {
        data() {
            return {
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
