import 'package:alpha_interface_app/screens/forward.dart';
import 'package:alpha_interface_app/screens/futures.dart';
import 'package:alpha_interface_app/screens/home.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';

class MyTabBar extends StatelessWidget {
  final String productName;
  final String currentPrice;
  final List<PriceData> data;
  MyTabBar(this.productName, this.currentPrice, this.data);

  @override
  Widget build(BuildContext context) {
    return DefaultTabController(
      initialIndex: 0,
      length: 2,
      child: Scaffold(
        appBar: AppBar(
          title: Center(child: Text(productName)),
          bottom: const TabBar(
            indicatorWeight: 10,
            indicatorColor: Colors.red,
            tabs: <Widget>[
              Tab(
                child: Center(child: Text("Futures Trading")),
              ),
              Tab(
                child: Center(child: Text("Forward Trading")),
              ),
            ],
          ),
        ),
        body: TabBarView(
          children: <Widget>[
            FutureTrading(productName, currentPrice),
            ForwardTrading(productName, currentPrice),
          ],
        ),
      ),
    );
  }
}
