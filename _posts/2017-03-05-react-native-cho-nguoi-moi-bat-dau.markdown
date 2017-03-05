---
layout: post
title: ReactNative Cho Người Mới Bắt Đầu
date: 2017-03-05 18:00
comments: true
external-url: 
categories: ReactNative
keywords: react, react native, reactnative, ios, android
excerpt: Chắc hẳn các bạn không còn xa lạ với React Native (RN) nên mình sẽ không giới thiệu nó là cái gì nữa mà sẽ bắt tay vào làm một app hello react native ngay :D
image:
  feature: ../../assets/deploy-your-rails-app-to-aws.png
  credit: 
  creditlink:
---
Chắc hẳn các bạn không còn xa lạ với React Native (RN) nên mình sẽ không giới thiệu nó là cái gì nữa mà sẽ bắt tay vào làm một app hello react native ngay :D

# Cài Đặt Môi Trường

Các bạn có thể lập trình trên Windows hay OSX hay Linux đều được nhưng tốt nhất là dùng OSX cho ít tốn thời gian và công sức .

Môi trường phát triển OSX:
- Máy ảo phát triển iOS
- Máy ảo phát triển Android

Môi trường phát triển Windows:
- Máy ảo phát triển iOS (Hiện tại không hỗ trợ)
- Máy ảo phát triển Android

Môi trường phát triển Linux:
- Máy ảo phát triển iOS (Hiện tại không hỗ trợ)
- Máy ảo phát triển Android

## Ở đây mình tập trung vào lập trình RN trên OSX :

`brew update`

- Cài đặt Node, Watchman :

```
brew install node

brew install watchman
```

[Watchman](https://facebook.github.io/watchman/) là một công cụ của Facebook để theo dõi sự thay đổi trong file hệ thống. Nó được khuyến khích cài đặt để nâng cao hiệu suất công việc của bạn.

React Native CLI :
Node.js đã cung cấp [npm](https://www.npmjs.com), và chúng ta sẽ sử dụng nó để cài đặt giao diện dòng lệnh cho React Native. Chạy lệnh sau ở cửa sổ Terminal

```
npm install -g react-native-cli
```

Xcode : Bạn có thể dễ dàng cài đặt Xcode thông qua Mac App Store. Việc cài đặt Xcode sẽ đồng thời cài đặt máy ảo iOS và tất cả những công cụ cần thiết để bạn có thể build ứng dụng iOS.

# Tạo ứng dụng Hello React Native

Sử dụng giao diện dòng lệnh React Native đẻ tạo ra một project mới của React Native ví dụ như HelloReactNative sau đó chạy lệnh react-native run-ios bên trong thư mục project mới được tạo.

```
react-native init HelloReactNative
cd HelloReactNative
react-native run-ios
```

Bạn sẽ nhìn thấy ứng dụng mới của bạn được chạy trên máy ảo iOS.
Câu lệnh react-native run-ios là cách để chạy ứng dụng của bạn. Bạn cũng có thể chạy ứng dụng này trong Xcode hoặc Nuclide.
Sau khi ứng dụng được mở lên ta sẽ nhìn thấy màn hình mặc định ban đầu của RN :

![](/assets/welcome-react-native.png){:height="50%" width="50%"}

## Bắt đầu chỉnh sửa app của bạn :

Bây giờ bạn đã chạy thành công ứng dụng vừa mới tạo. Hãy chỉnh sửa nó

- Mở file `index.ios.js` bằng bất kỳ trình sửa text nào và chỉnh sửa nội dung trong đó.
Dùng tổ hợp phím **Command⌘ + R** trong máy ảo iOS để reload sự thay đổi của ứng dụng sau khi chỉnh sửa.(Đây chính là cái hay nhất của RN k cần phải build lại nhưng vẫn sẽ cập nhật các thay đổi)

- Ta sẽ tạo 1 file mới HelloWorld.js:

```
'use strict';

  import React, { Component } from 'react';
  import {
    StyleSheet,
    Text,
    View
  } from 'react-native';
  class HelloWorld extends Component {
    render() {
      return (
        <View style={styles.container}>
          <Text style={styles.welcome}>
            Welcome to React Native!
          </Text>
          <Text style={styles.instructions}>
            To get started, edit index.ios.js
          </Text>
          <Text style={styles.instructions}>
            Press Cmd+R to reload,{'\n'}
            Cmd+D or shake for dev menu
          </Text>
        </View>
      );
    }
  }

  const styles = StyleSheet.create({
    container: {
      flex: 1,
      justifyContent: 'center',
      alignItems: 'center',
      backgroundColor: '#F5FCFF',
    },
    welcome: {
      fontSize: 20,
      textAlign: 'center',
      margin: 10,
    },
    instructions: {
      textAlign: 'center',
      color: '#333333',
      marginBottom: 5,
    },
  });

  export default HelloWorld;
```

`'use strict';` : Dòng này kích hoạt chế độ Strict Mode, nó tăng cường khả năng xử lí lỗi của Javascript.

```
render() {
          return (
            <View style={styles.container}>
              <Text style={styles.welcome}>
                Welcome to React Native!
              </Text>
              <Text style={styles.instructions}>
                To get started, edit index.ios.js
              </Text>
              <Text style={styles.instructions}>
                Press Cmd+R to reload,{'\n'}
                Cmd+D or shake for dev menu
              </Text>
            </View>
          );
        }
```

Đoạn này sẽ tạo ra một class chỉ có function duy nhất là render. Hàm render sẽ return lại những gì sẽ được hiển thị lên màn hình. Đoạn code trong phần return sử dụng JSX (Javascript syntax extension). Nếu bạn đã làm việc với React.JS thì đoạn code trên rất quen thuộc.

```
const styles = StyleSheet.create({
        container: {
          flex: 1,
          justifyContent: 'center',
          alignItems: 'center',
          backgroundColor: '#F5FCFF',
        },
        welcome: {
          fontSize: 20,
          textAlign: 'center',
          margin: 10,
        },
        instructions: {
          textAlign: 'center',
          color: '#333333',
          marginBottom: 5,
        },
      });
```

Đoạn trên là style dùng cho code JSX bên trên. Đoạn code này rất quen thuộc với những ai làm web vì React Native sử dụng CSS để làm style cho giao diện app. Nếu bạn nhìn lên code JSX ở trên thì bạn sẽ thấy cách mỗi style được sử dụng. Ví dụ component View có style={styles.container} thì các định nghĩa về giao diện của container sẽ được dùng cho View.

và chỉnh sửa một tí file `index.ios.js`

```
'use strict';

        import React, { Component } from 'react';
        import {
          AppRegistry,
          StyleSheet,
          Text,
          View,
          NavigatorIOS
        } from 'react-native';

        import HelloWorld from './HelloWorld';

        class HelloReact extends Component {
          render() {
            return (
              <NavigatorIOS
                style={styles.container}
                initialRoute= { { 
                  title: 'Hello React App',
                  component: HelloWorld
                }} />
            );
          }
        }

        const styles = StyleSheet.create({
          text: {
            color: 'black',
            backgroundColor: 'white',
            fontSize: 30,
            margin: 80
          },
          container: {
            flex: 1
          }
        });

        export default HelloReact;

        AppRegistry.registerComponent('HelloReact', () => HelloReact
```

```
import {
    AppRegistry,
    StyleSheet,
    Text,
    View,
    NavigatorIOS
} from 'react-native';
```

Đoạn này sẽ load module react-native, gán vào biến React. Đồng thời gán các thuộc tính của React như `React.AppRegistry`, `React.StyleSheet` vào các biến tương tự cùng tên. Điều này giúp bạn viết code ngắn gọn hơn, ví dụ viết `AppRegistry` thay vì `React.AppRegistry`.

`NavigatorIOS` ở đây sẽ có nhiệm vụ tạo ra một navigation controller, và gán route mặc định cho nó là component HelloWorld, đồng thời đặt tiêu đề sẽ hiển thị trên navigation bar.

```
import HelloWorld from './HelloWorld';
```

Đoạn này load class `HelloWorld` từ file `HelloWorld.js` để sử dụng.

```
AppRegistry.registerComponent('HelloReact', () => HelloReact
```

Đoạn này định nghĩa điểm khởi đầu cho chương trình, nơi mà Javascript bắt đầu thực thi.

Đó là cấu trúc cơ bản của React Native UI. Sau này, tất cả các view chúng ta làm ra đều tuân theo cấu trúc cơ bản như vậy.

Reload lại app, và chúng ta được như sau:

![](/assets/hello-react-native.png){:height="50%" width="50%"}

Các bạn có thể tham khảo thêm để làm 1 app RN đơn giản tại [đây](https://www.raywenderlich.com/126063/react-native-tutorial)

Trang chủ RN : [React Native](https://facebook.github.io/react-native/docs/getting-started.html)

Source code ví dụ này : [HelloReactNative](https://github.com/minhquan4080/HelloReactNative)

# Một số tài liệu:

[React Native Tutorial](https://www.tutorialspoint.com/react_native/index.htm)

[Let's build a React Native app in 20 minutes - YouTube](https://www.youtube.com/watch?v=9ArhJiMGVDc)

[React Native Fundamentals - Course by @tylermcginnis33 ](https://egghead.io/courses/react-native-fundamentals)

[Build iOS Apps with React Native - Pluralsight](https://www.pluralsight.com/courses/build-ios-apps-react-native)

[The Complete React Native and Redux Course - Udemy](https://www.udemy.com/the-complete-react-native-and-redux-course/?altsc=610300)

[Build Cross Platform React Native Apps with Exponent and Redux - Pluralsight](https://www.pluralsight.com/courses/build-react-native-exponent-redux-apps)

[react-native training · GitBook](https://www.gitbook.com/book/unbug/react-native-training/details)

[Learn React Native Quickly](https://www.dailydrip.com/topics/react-native/)

[Use this to learn React Native](https://github.com/hsavit1/Awesome-React-Native-Education )

[A Complete Guide to Flexbox .](https://css-tricks.com/snippets/css/a-guide-to-flexbox/ )

[GitHub - ericdouglas/ES6-Learning: List of resources to learn ECMAScript 6!](https://github.com/ericdouglas/ES6-Learning )

[GitHub - happypoulp/redux-tutorial: Learn how to use redux step by step](https://github.com/happypoulp/redux-tutorial)

[Mac OS X Development Tutorial for Beginners Part 1: Intro to Xcode](https://www.raywenderlich.com/110170/mac-os-x-development-tutorial-for-beginners-part-1-intro-to-xcode)

[Android Studio](https://www.tutorialspoint.com/android/android_studio.htm )

[40 Terminal Tips and Tricks You Never Thought You Needed](https://computers.tutsplus.com/tutorials/40-terminal-tips-and-tricks-you-never-thought-you-needed--mac-51192 )

 [CocoaPods Guides - Using CocoaPods](https://guides.cocoapods.org/using/using-cocoapods.html)
 
### Nên dùng :
 [Native Base](http://nativebase.io) giống như 1 framework cho lập trình RN , hỗ trợ đa dạng các components, đơn giản dể hiểu dễ sửa.
 
 React Native Starter App with NativeBase + CodePush + Redux :  [GitHub - start-react/native-starter-kit](https://github.com/start-react/native-starter-kit) : 1 project mẫu thích hợp để bắt đầu.
 
### Awesome React Native

 [GitHub - jondot/awesome-react-native](https://github.com/jondot/awesome-react-native ) : Tổng hợp các module hữu ích cho RN mà mọi developer đều cần :D


