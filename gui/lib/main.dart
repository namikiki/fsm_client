import 'dart:convert';
import 'dart:io';
import 'package:http/http.dart' as http;
import 'package:flutter/material.dart';
import 'login.dart';
import 'setting.dart';
import 'sync_task_list.dart';
import 'consts.dart';
import 'package:window_manager/window_manager.dart';

void main() async {
  WidgetsFlutterBinding.ensureInitialized();
  // Must add this line.
  await windowManager.ensureInitialized();

  WindowOptions windowOptions = const WindowOptions(
    size: Size(800, 600),
    center: true,
    backgroundColor: Colors.transparent,
    // skipTaskbar: true,
    titleBarStyle: TitleBarStyle.normal,
  );
  windowManager.waitUntilReadyToShow(windowOptions, () async {
    await windowManager.show();
    await windowManager.focus();
  });

  runApp(const MyApp());
}

class MyApp extends StatefulWidget {
  const MyApp({super.key});

  @override
  MyAppState createState() => MyAppState();
}

class MyAppState extends State<MyApp> with WindowListener {
  late Process _process;

  Future<void> _startProcess() async {
    // 等待进程完成
    String binaryPath = goBinPath;
    _process = await Process.start(binaryPath, []);

    _process.stdout.transform(utf8.decoder).listen((data) {
      print(data);
    });

    // 获取二进制文件的错误输出
    _process.stderr.transform(utf8.decoder).listen((data) {
      print(data);
    });
  }

  Future<void> _stopProcess() async {
    _process.kill(ProcessSignal.sigterm);
  }

  @override
  void initState() {
    windowManager.addListener(this);
    super.initState();
    _startProcess();
  }

  @override
  void dispose() {
    windowManager.removeListener(this);
    super.dispose();
  }

  @override
  void onWindowClose() {
    _stopProcess();
  }

  @override
  Widget build(BuildContext context) {
    return const MaterialApp(
      home: HomePage(),
    );
  }
}

class HomePage extends StatefulWidget {
  const HomePage({super.key});

  @override
  HomePageState createState() => HomePageState();
}

class HomePageState extends State<HomePage> {
  bool _isLoading = true;
  bool _isError = false;
  String w = '';
  final _addFormKey = GlobalKey<FormState>();
  String _name = '';
  String _path = '';
  bool _ignore = true;
  bool _isSwitched = false;
  late Widget _myWidget;

  @override
  void didUpdateWidget(HomePage oldWidget) {
    super.didUpdateWidget(oldWidget);
  }

  @override
  void initState() {
    super.initState();
    _loginState();
  }

  Future<void> _loginState() async {
    for (var i = 0; i < 3; i++) {
      var getStatusURL = Uri.http(serverAddress, "status");
      var response = await http.get(getStatusURL);

      if (response.statusCode == 200) {
        _myWidget = const DynamicListTable();
        setState(() {
          _isLoading = false;
        });
      } else if (response.statusCode == 502) {
        setState(() {
          _isLoading = false;
          _isError = true;
        });
        break;
      } else {
        await Future.delayed(const Duration(seconds: 1));
      }
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('文件同步管理器'),
        actions: <Widget>[
          _isError
              ? const SizedBox()
              : IconButton(
                  icon: const Icon(Icons.add),
                  onPressed: () {
                    showDialog(
                      context: context,
                      builder: (BuildContext context) {
                        return AlertDialog(
                          title: const Text('添加同步任务'),
                          content: Form(
                            key: _addFormKey,
                            child: Column(
                              mainAxisSize: MainAxisSize.min,
                              children: <Widget>[
                                Row(
                                  children: [
                                    const Text("同步名"),
                                    Expanded(
                                        child: TextFormField(
                                          decoration: const InputDecoration(
                                            contentPadding: EdgeInsets.fromLTRB(
                                                10, 0, 0, 0), // 设置内边距为零
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
                                          decoration: const InputDecoration(
                                            contentPadding: EdgeInsets.fromLTRB(
                                                10, 0, 0, 0), // 设置内边距为零
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
                                          //重新构建页面
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
                              padding: const EdgeInsets.symmetric(
                                  horizontal: 16, vertical: 12),
                              child: Row(
                                mainAxisAlignment: MainAxisAlignment.spaceBetween,
                                children: [
                                  OutlinedButton(
                                    child: Text("取消"),
                                    onPressed: () {},
                                  ),
                                  OutlinedButton(
                                    child: Text("同步"),
                                    onPressed: () async {
                                      final addForm = _addFormKey.currentState;
                                      addForm?.save();

                                      var url = Uri.http(serverAddress, "syncTask");
                                      var response = await http.post(url,
                                          body: jsonEncode({
                                            'name': _name,
                                            'path': _path,
                                            'type': "two",
                                            'ignore': _ignore
                                          }));
                                      // print(response.statusCode);
                                      // if(response.statusCode == 200){
                                      // Navigator.of(context).pushReplacement(
                                      // MaterialPageRoute(builder: (context) => HomePage()),
                                      // );
                                      // await Future.delayed(const Duration(seconds: 1));
                                      // Navigator.of(context).pop();
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
                  },
                ),

          const IconButton(onPressed: null, icon: Icon(Icons.sync)),
          const Padding(padding: EdgeInsets.only(right: 16)),
        ],
      ),
      drawer: LayoutBuilder(
        builder: (BuildContext context, BoxConstraints constraints) {
          return SizedBox(
            width: constraints.maxWidth * 0.2 > 150
                ? constraints.maxWidth * 0.2
                : 150, // set the width of the drawer to 70% of the screen width
            child: Drawer(
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.stretch,
                children: <Widget>[
                  const SizedBox(
                    height: 100,
                    child: DrawerHeader(
                      decoration: BoxDecoration(
                        color: Colors.blue,
                      ),
                      child: Text(
                        '用户名',
                        style: TextStyle(
                          color: Colors.white,
                          fontSize: 24,
                        ),
                      ),
                    ),
                  ),
                  ListTile(
                    leading: Icon(Icons.home),
                    title: const Text('主页'),
                    onTap: () {
                      Navigator.pop(context);
                    },
                  ),
                  ListTile(
                    leading: Icon(Icons.settings),
                    title: const Text('设置'),
                    onTap: () {
                      Navigator.pop(context);
                      Navigator.push(
                        context,
                        MaterialPageRoute(
                          builder: (context) => SettingsPage(),
                        ),
                      );
                    },
                  ),
                ],
              ),
            ),
          );
        },
      ),
      body: _isLoading
          ? const Center(
              child: CircularProgressIndicator(),
            )
          : _isError
              ? const LoginForm()
              : _myWidget,
    );
  }
}

class MyHomePage extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('My App'),
      ),
      body: Center(
        child: IconButton(
          icon: Icon(Icons.add),
          onPressed: () {
            Navigator.of(context).push(
              MaterialPageRoute(
                builder: (BuildContext context) => MyDialog(),
                fullscreenDialog: true,
              ),
            );
          },
        ),
      ),
    );
  }
}

class MyDialog extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return AlertDialog(
      title: Text('My Dialog'),
      content: Text('This is my dialog.'),
      actions: <Widget>[
        TextButton(
          onPressed: () {
            Navigator.of(context).pop();
          },
          child: const Text('取消'),
        ),
        TextButton(
            child: const Text('确定'),
            onPressed: () async {
              var url = Uri.http("baidu.com");
              var response = await http.get(url);

              if (response.statusCode == 200) {
                Navigator.of(context).pop();
              }
            }),
      ],
    );
  }
}
