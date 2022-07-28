# 存储卷管理



## 统一状态


### 各云商状态

阿里云: [云盘状态表](https://help.aliyun.com/document_detail/25689.html?spm=api-workbench.API%20Explorer.0.0.40731e0ffIqygN)
+ Creating	创建中。 通过RunInstances、CreateInstance或CreateDisk创建了云盘后，云盘进入短暂的创建中状态。
+ Available	待挂载。 通过CreateDisk成功创建一块按量付费云盘或通过DetachDisk卸载一块按量付费数据盘后，云盘进入稳定的待挂载状态。
+ In_Use	使用中。 云盘的稳定状态，
+ ReIniting	初始化中。通过ReInitDisk重新初始化一块系统盘或者数据盘后，云盘进入短暂的初始化中状态。
+ Detaching	卸载中。 通过DetachDisk卸载一块按量付费数据盘后，云盘进入短暂的卸载中状态。
+ Deleting*	删除中。 通过DeleteDisk释放一块按量付费数据盘后，云盘进入短暂的删除中状态。
+ Deleted*	已删除。 通过DeleteDisk释放一块按量付费数据盘后，云盘进入短暂的已删除状态。

腾讯云: [Disk数据结构](https://cloud.tencent.com/document/api/362/15669#Disk) 云盘状态。取值范围：
+ UNATTACHED：未挂载
+ ATTACHING：挂载中
+ ATTACHED：已挂载
+ DETACHING：解挂中
+ EXPANDING：扩容中
+ ROLLBACKING：回滚中
+ TORECYCLE：待回收
+ DUMPING：拷贝硬盘中

华为云: [云硬盘状态](https://support.huaweicloud.com/api-evs/evs_04_0040.html)
+ creating 云硬盘处于正在创建的过程中。
+ available 云硬盘创建成功，还未挂载给任何云服务器，可以进行挂载。
+ in-use 云硬盘已挂载给云服务器，正在使用中。
+ error 云硬盘在创建过程中出现错误。
+ attaching 云硬盘处于正在挂载的过程中。
+ detaching 云硬盘处于正在卸载的过程中。
+ restoring-backup 云硬盘处于正在从备份恢复的过程中。
+ backing-up 云硬盘处于通过备份创建的过程中。
+ error_restoring 云硬盘从备份恢复过程中出现错误。
+ uploading 云硬盘数据正在被上传到镜像中。此状态出现在从云服务器创建镜像的操作过程中。
+ downloading 正在从镜像下载数据到云硬盘。此状态出现在创建云服务器的操作过程中。
+ extending 云硬盘处于正在扩容的过程中。
+ error_extending 云硬盘在扩容过程中出现错误。
+ deleting 云硬盘处于正在删除的过程中。
+ error_deleting云硬盘在删除过程中出现错误。
+ rollbacking 云硬盘处于正在从快照回滚数据的过程中。
+ error_rollbacking 云硬盘在从快照回滚数据的过程中出现错误。
+ awaiting-transfer 云硬盘处于等待过户的过程中。

### 统一后状态

以腾讯云状态作为统一: 
+ UNATTACHED：未挂载
+ ATTACHING：挂载中
+ ATTACHED：已挂载
+ DETACHING：解挂中
+ EXPANDING：扩容中
+ ROLLBACKING：回滚中
+ TORECYCLE：待回收
+ DUMPING：拷贝硬盘中