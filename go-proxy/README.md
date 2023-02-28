使用 go 代理

### 案例

采集 TRON 数据

## 数据

### 速率



官方提示请求不超 25个/s，实际请求 16个/s。

用代理测试 9个/s。约等于 77.76w

每天请求上限：官方提示10w

{"Error":"request rate exceeded the allowed_rps(25), and the query server is suspended for 1 s"}


### 

代理能否增加请求上限？

可以。



### 区别不同交易的交易类型


"type": "TransferContract" trx、
"type": "TransferAssetContract" trc10、
"type": "TriggerSmartContract" 触发合约：USDT、
"type": "FreezeBalanceContract" 质押资产、冻结资产、
"type": "UnfreezeBalanceContract" 解锁资产、解冻资产、
"type": "VoteWitnessContract" 投票、
"type": "WithdrawBalanceContract" 领取投票收益、
"type": "AccountPermissionUpdateContract" 更新账户权限、
"type": "AccountCreateContract" 激活账户、
"type": "AccountUpdateContract" 修改账户名称、
"type": "ParticipateAssetIssueContract" 购买TRC10通证
"type": "AssetIssueContract" 发行TRC10通证
"type": "UnfreezeAssetContract" TRC10锁仓提取
"type": "UpdateAssetContract" 更新TRC10通证参数
"type": "CreateSmartContract" 创建智能合约
"type": "UpdateSettingContract" 更新合约参数
"type": "UpdateEnergyLimitContract" 更新合约能量限制
"type": "ClearABIContract" 清除合约ABI
"type": "ProposalCreateContract" 超级代表-创建提议
"type": "ProposalApproveContract" 赞成提议
"type": "ProposalDeleteContract" 撤销提议
"type": "WitnessCreateContract" 创建超级代表候选人
"type": "WitnessUpdateContract" 更新超级代表候选人信息
"type": "UpdateBrokerageContract" 更新超级代表佣金比例 
"type": "ExchangeTransactionContract" 执行Bancor交易
"type": "ExchangeCreateContract" 创建Bancor交易
"type": "ExchangeInjectContract" Bancor交易注资
"type": "ExchangeWithdrawContract" Bancor交易撤资



Tron 交易类型：https://zhuanlan.zhihu.com/p/110595685


