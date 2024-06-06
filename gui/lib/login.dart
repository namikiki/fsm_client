import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import 'consts.dart';
import 'dart:convert';
import 'main.dart';

class LoginForm extends StatefulWidget {
  const LoginForm({Key? key}) : super(key: key);

  @override
  LoginFormState createState() => LoginFormState();
}

class LoginFormState extends State<LoginForm> {
  final TextEditingController _emailController = TextEditingController();
  final TextEditingController _passwordController = TextEditingController();
  String _loginText = "Login";

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: SafeArea(
        child: Center(
          child: SingleChildScrollView(
            padding: const EdgeInsets.symmetric(horizontal: 24),
            child: Column(
              mainAxisAlignment: MainAxisAlignment.center,
              children: [
                // Image.asset(
                //   'assets/logo.png',
                //   height: 120,
                // ),
                const SizedBox(height: 32),
                _buildEmailInput(),
                const SizedBox(height: 16),
                _buildPasswordInput(),
                const SizedBox(height: 16),
                ElevatedButton(
                    style:ElevatedButton.styleFrom(
                      minimumSize: const Size(200, 50),
                    ),
                    onPressed: () async {
                      final String email = _emailController.text.trim();
                      final String password = _passwordController.text.trim();

                      if (email.isEmpty) {
                        ScaffoldMessenger.of(context)
                            .showSnackBar(const SnackBar(
                          content: Text('请输入你注册的邮箱'),
                        ));
                        return;
                      }

                      if (password.isEmpty) {
                        ScaffoldMessenger.of(context)
                            .showSnackBar(const SnackBar(
                          content: Text('请输入你的密码'),
                        ));
                        return;
                      }

                      var url = Uri.http(serverAddress, "login");
                      var response = await http.post(url,
                          body: jsonEncode(
                              {'email': email, 'password': password}));
                      print(response.statusCode);

                      if (response.statusCode == 200) {
                        Navigator.of(context).pushReplacement(
                          MaterialPageRoute(builder: (context) => HomePage()),
                        );
                      } else {
                        setState(() {
                          _loginText = "login err";
                        });
                      }

                      //   }
                      //   } catch (e) {
                      //   setState(() {
                      //   _errorMessage = e.toString();
                      //   print(_errorMessage);
                      //   });
                      //   }
                      // }


                    },
                    child: Text(_loginText)),
              ],
            ),
          ),
        ),
      ),
    );
  }

  Widget _buildEmailInput() {
    return SizedBox(
      width: 400,
      child: TextField(
        controller: _emailController,
        keyboardType: TextInputType.emailAddress,
        decoration: const InputDecoration(
          labelText: 'Email',
          border: OutlineInputBorder(),
          contentPadding: EdgeInsets.symmetric(vertical: 12, horizontal: 16),
          prefixIcon: Icon(Icons.email),
        ),
      ),
    );
  }

  Widget _buildPasswordInput() {
    return SizedBox(
        width: 400,
        child: TextField(
          controller: _passwordController,
          obscureText: true,
          decoration: const InputDecoration(
            labelText: 'Password',
            border: OutlineInputBorder(),
            contentPadding: EdgeInsets.symmetric(vertical: 12, horizontal: 16),
            prefixIcon: Icon(Icons.lock),
            suffixIcon: Icon(Icons.visibility),
          ),
        ));
  }
}

// class LoginForm extends StatefulWidget {
//    const LoginForm({Key? key}) : super(key: key);
//
//   @override
//   LoginFormState createState() => LoginFormState();
// }
//
// class LoginFormState extends State<LoginForm> {
//   final _formKey = GlobalKey<FormState>();
//
//   String _email = '';
//   String _password = '';
//   String _errorMessage = '';
//   String _loginText = "登录";
//
//
//
//
//
// // // TODO: 注册和登录方法
// //   @override
// //   Widget build(BuildContext context) {
// //     return Form(
// //         key: _formKey,
// //         child: Padding(
// //           padding: EdgeInsets.all(80),
// //           child: Column(
// //             mainAxisSize: MainAxisSize.min,
// //             children: [
// //               TextFormField(
// //                 keyboardType: TextInputType.emailAddress,
// //                 decoration: const InputDecoration(labelText: '邮箱'),
// //                 validator: (value) {
// //                   if (value == null || value.isEmpty || !value.contains('@')  ) {
// //                     return 'Please enter a valid email address.';
// //                   }
// //                   return null;
// //                 },
// //                 onChanged: (value){
// //
// //                    if (_loginText !="login"){
// //                       setState(() {
// //                         _loginText = "login";
// //                       });
// //                    }
// //
// //                 },
// //                 onSaved: (value) {
// //                   _email = value!;
// //                 },
// //               ),
// //               TextFormField(
// //                 obscureText: true,
// //                 decoration: const InputDecoration(labelText: '密码'),
// //                 validator: (value) {
// //                   if (value == null || value.isEmpty || value.length < 7) {
// //                     return 'Password must be at least 7 characters long.';
// //                   }
// //
// //                   return null;
// //                 },
// //                 onChanged: (value){
// //
// //                   if (_loginText !="login"){
// //                     setState(() {
// //                       _loginText = "login";
// //                     });
// //                   }
// //
// //                 },
// //                 onSaved: (value) {
// //                   _password = value!;
// //                 },
// //               ),
// //               const SizedBox(height: 12),
// //               ElevatedButton(
// //                 onPressed: onSubmit,
// //                 child:   Text(_loginText),
// //               ),
// //             ],
// //           ),
// //         ));
// //   }
// //
// //   void onSubmit() async {
// //     final form = _formKey.currentState;
// //     if (form!.validate()) {
// //       form.save();
// //       try {
// //         print(_email);
// //         print(_password);
// //         print(_formKey);
// //
// //         var url = Uri.http(serverAddress,"login");
// //         var response = await http.post(url, body: jsonEncode({'email': _email, 'password': _password}));
// //         print(response.statusCode);
// //         if(response.statusCode == 200){
// //           Navigator.of(context).pushReplacement(
// //             MaterialPageRoute(builder: (context) =>  HomePage()),
// //           );
// //         } else {
// //
// //           setState(() {
// //             _loginText = "login err";
// //
// //           });
// //         }
// //       } catch (e) {
// //         setState(() {
// //           _errorMessage = e.toString();
// //           print(_errorMessage);
// //         });
// //       }
// //     }
// //   }
//
//   Future<void> register() async {
//     //   final auth = FirebaseAuth.instance;
//     //   try {
//     //     final userCredential = await auth.createUserWithEmailAndPassword(
//     //       email: _email,
//     //       password: _password,
//     //     );
//     //   final
//     //
//     // }
//     // }
//   }
//
// }
