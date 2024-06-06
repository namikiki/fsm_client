import 'dart:async';
import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:fsm_gui/consts.dart';
import 'package:http/http.dart' as http;

class DynamicListTable extends StatefulWidget {
  const DynamicListTable({super.key});

  @override
  State<DynamicListTable> createState() => _DynamicListTableState();
}

class SyncTask {
  final String id;
  final String name;
  final String rootDir;
  late String operate;
  final String status;

  SyncTask({
    required this.id,
    required this.name,
    required this.rootDir,
    required this.status,
    required this.operate,
  });

  // factory SyncTask.fromJson(Map<String, dynamic> parsedJson) {
  //   return SyncTask(
  //
  //       id: parsedJson['id'],
  //       name: parsedJson['name'],
  //       rootDir: parsedJson['root_dir'], status: '');
  // }
}

class _DynamicListTableState extends State<DynamicListTable> {
  List<SyncTask> syncList = <SyncTask>[];

  Map<String, IconData> iconMap = {
    "created": Icons.add_circle_outline,
    "syncing": Icons.sync_outlined,
    "sync": Icons.check_circle_outline,
    'pause': Icons.sync_lock,
    "update": Icons.cloud_sync,
    "delete": Icons.sync_disabled
  };

  Map<String, IconData> opIcon = {
    "syncing": Icons.pause,
    'pause': Icons.play_arrow,
    "sync": Icons.pause,
    "update": Icons.pause,
    'created': Icons.backup,
  };

  final _formKeys = GlobalKey<FormState>();
  String _name = '';
  String _path = '';
  bool _ignore = true;

  void loopGetTask() {
    Timer.periodic(const Duration(seconds: 2), (timer) {
      GetSyncTask();
    });
  }

  Future<void> GetSyncTask() async {
    var url = Uri.http(serverAddress, "syncTasks");
    var response = await http.get(url);
    List<dynamic> taskList = jsonDecode(response.body);
    List<SyncTask> cache = <SyncTask>[];

    for (var e in taskList) {
      var st = SyncTask(
          id: e['id'],
          name: e['name'],
          rootDir: e['root_dir'],
          status: e['status'],
          operate: e['status']);

      if (e['status'] == "delete") {
        st.operate = "123";
      }
      cache.add(st);
    }

    setState(() {
      syncList = cache;
    });

    // taskList.map((task) =>  print(task));\
  }

  @override
  void initState() {
    super.initState();
    GetSyncTask();
    loopGetTask();
  }

  Future<void> pauseSync(int index) async {
    if (syncList[index].status == "created") {
      _recoverTask(index);
      return;
    }

    var url = Uri.http(serverAddress, "pause");
    var response =
        await http.post(url, body: jsonEncode({'id': syncList[index].id}));
    GetSyncTask();
  }

  //Recover syncTask from cloud to local alertDialog
  void _recoverTask(int index) {
    showDialog(
      context: context,
      builder: (BuildContext context) {
        return AlertDialog(
          title: const Text('恢复同步任务'),
          content: Form(
            key: _formKeys,
            child: Column(
              mainAxisSize: MainAxisSize.min,
              children: <Widget>[
                Row(
                  children: [
                    Text("同步名"),
                    Expanded(
                        child: TextFormField(
                      initialValue: syncList[index].name,
                      decoration: const InputDecoration(
                        contentPadding:
                            EdgeInsets.fromLTRB(10, 0, 0, 0), // 设置内边距为零
                        isCollapsed: true, // 设置输入框内部大小为实际内容大小
                        border: InputBorder.none,
                      ),
                      obscureText: false,
                      onSaved: (value) {
                        _name = value!;
                      },
                    ))
                  ],
                ),
                Row(
                  children: [
                    Text("同步路径"),
                    Expanded(
                        child: TextFormField(
                      initialValue: syncList[index].rootDir,
                      decoration: const InputDecoration(
                        contentPadding:
                            EdgeInsets.fromLTRB(10, 0, 0, 0), // 设置内边距为零
                        isCollapsed: true, // 设置输入框内部大小为实际内容大小
                        border: InputBorder.none,
                      ),
                      obscureText: false,
                      onSaved: (value) {
                        _path = value!;
                      },
                    )),
                  ],
                ),
                Row(
                  children: [
                    const Text("启用过滤"),
                    Expanded(
                      child: Switch(
                        value: _ignore, //当前状态
                        onChanged: (value) {
                          setState(() {
                            _ignore = value;
                          });
                        },
                      ),
                    )
                  ],
                ),
              ],
            ),
          ),
          actions: <Widget>[
            Padding(
              //上下左右各添加16像素补白
              padding: const EdgeInsets.symmetric(horizontal: 16, vertical: 12),
              child: Row(
                mainAxisAlignment: MainAxisAlignment.spaceBetween,
                children: [
                  OutlinedButton(
                    child: Text("取消"),
                    onPressed: () {
                      Navigator.pop(context);
                    },
                  ),
                  OutlinedButton(
                    child: Text("恢复"),
                    onPressed: () async {
                      final form = _formKeys.currentState;
                      form?.save();

                      var url = Uri.http(serverAddress, "recover");
                      var response = await http.post(url,
                          body: jsonEncode({
                            'id': syncList[index].id,
                            'name': _name,
                            'path': _path,
                            'ignore': _ignore
                          }));
                      GetSyncTask();

                      if (response.statusCode == 200) {
                        Navigator.pop(context);
                      }
                    },
                  ),
                ],
              ),
            )
          ],
        );
      },
    );
  }

  //Delete syncTask button function
  Future<void> deleteTask(int index) async {
    var url = Uri.http(serverAddress, "syncTask");
    var response = await http.delete(url,
        body: jsonEncode({
          'id': syncList[index].id,
          'del_local': false,
          'del_cloud': false,
        }));
    GetSyncTask();
    print("删除同步任务结果${response.statusCode}");
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: ListView.builder(
        itemCount: syncList.length,
        itemExtent: 50,
        itemBuilder: (BuildContext context, int index) {
          return Container(
            margin: const EdgeInsets.all(4.0),
            padding: const EdgeInsets.fromLTRB(16, 4, 16, 4),
            decoration: BoxDecoration(
              borderRadius: BorderRadius.circular(4.0),
              border: Border.all(
                  color: Colors.black45, width: 1.0, style: BorderStyle.solid),
            ),
            child: Flex(
              // mainAxisAlignment: MainAxisAlignment.spaceBetween,
              direction: Axis.horizontal,
              children: <Widget>[
                Expanded(
                    flex: 1,
                    child: Text(syncList[index].name,
                        style: const TextStyle(fontSize: 12.0))),
                Expanded(
                    flex: 2,
                    child: Row(
                        mainAxisAlignment: MainAxisAlignment.start,
                        children: <Widget>[
                          const Icon(Icons.folder),
                          Text(
                            syncList[index].rootDir,
                            style: const TextStyle(fontSize: 12.0),
                          ),
                        ])),
                Expanded(
                    flex: 1,
                    child: Row(
                      mainAxisAlignment: MainAxisAlignment.spaceBetween,
                      children: <Widget>[
                        Icon(iconMap[syncList[index].status]),
                        InkWell(
                          borderRadius: BorderRadius.circular(50),
                          onTap: () {},
                          child: const Icon(
                            Icons.open_in_new_outlined,
                          ),
                        ),
                        InkWell(
                          borderRadius: BorderRadius.circular(50),
                          onTap: () => pauseSync(index),
                          child: Icon(
                            opIcon[syncList[index].operate],
                          ),
                        ),
                        InkWell(
                          borderRadius: BorderRadius.circular(50),
                          onTap: () => deleteTask(index),
                          child: const Icon(
                            Icons.delete,
                          ),
                        ),
                        // IconButton(onPressed: null, icon: Icon(Icons.delete_forever)),
                      ],
                    ))
              ],
            ),
          );
        },
      ),
    );
  }
}
