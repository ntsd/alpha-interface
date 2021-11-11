import 'package:alpha_interface_app/api/jsonAPI.dart';
import 'package:alpha_interface_app/model/yield.dart';
import 'package:alpha_interface_app/screens/home.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:syncfusion_flutter_charts/charts.dart';

class ForwardTrading extends StatefulWidget {
  final String productName;
  final String currentPrice;
  final List<PriceData> data;
  ForwardTrading(this.productName, this.currentPrice, this.data);

  @override
  _ForwardTradingState createState() => _ForwardTradingState();
}

class _ForwardTradingState extends State<ForwardTrading> {
  JsonAPI _jsonAPI = JsonAPI.getInstance();
  List<Yield> yieldList = [];
  String currentCountry = "";
  List<_SalesData> data = [];
  List<String> cities = [];
  int availableAmount = 1000;
  final TextEditingController _controller = TextEditingController();
  final TextEditingController _sellController = TextEditingController();
  int buyAmount = 0;
  int sellAmount = 0;

  @override
  void initState() {
    // TODO: implement initState
    super.initState();
    loadJson();
    _controller.text = "0";
    _sellController.text = "0";
  }

  loadJson() async {
    yieldList = await _jsonAPI.loadJson();
    currentCountry = yieldList.first.area;
    yieldList.forEach((element) {
      if (!cities.contains(element.area)) {
        cities.add(element.area);
      }
    });
    createChartDataForCountry();
    setState(() {});
  }

  createChartDataForCountry() {
    data = [];
    List<_SalesData> newDataList = [];
    yieldList.forEach((element) {
      if (element.area == currentCountry) {
        print(element.area +
            " " +
            element.year.toString() +
            " : " +
            element.value.toString());
        _SalesData newData = _SalesData(element.year, element.value);
        newDataList.add(newData);
      }
    });

    data.addAll(newDataList.where((v) => data.every((d) => v.year != d.year)));
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: SingleChildScrollView(
        child: Container(
          width: MediaQuery.of(context).size.width,
          child: Padding(
            padding: const EdgeInsets.all(35),
            child: Column(
              children: [
                Row(
                  mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                  children: [
                    Flexible(
                      flex: 2,
                      child: Padding(
                        padding: const EdgeInsets.all(10.0),
                        child: Container(
                          height: 50,
                          decoration: BoxDecoration(
                            color: Colors.white,
                            border: Border.all(width: 1),
                            borderRadius: BorderRadius.circular(2),
                          ),
                          child: Center(
                              child: Text(
                            "Germany",
                            style: TextStyle(fontSize: 30, color: Colors.black),
                          )),
                        ),
                      ),
                    ),
                    Flexible(
                      flex: 3,
                      child: Padding(
                        padding: const EdgeInsets.all(10.0),
                        child: Container(
                          decoration: BoxDecoration(
                            color: Colors.white,
                            border: Border.all(width: 1),
                            borderRadius: BorderRadius.circular(2),
                          ),
                          child: DropdownButtonHideUnderline(
                            child: ButtonTheme(
                              alignedDropdown: true,
                              child: DropdownButton<String>(
                                iconEnabledColor: Colors.black,
                                value: currentCountry,
                                isExpanded: true,
                                underline: Container(),
                                onChanged: (String value) {
                                  setState(() {
                                    currentCountry = value;
                                    createChartDataForCountry();
                                  });
                                },
                                items: createDropDownItems(),
                              ),
                            ),
                          ),
                        ),
                      ),
                    ),
                  ],
                ),
                Padding(
                  padding: const EdgeInsets.all(10.0),
                  child: SfCartesianChart(
                      primaryXAxis: CategoryAxis(),
                      tooltipBehavior: TooltipBehavior(enable: true),
                      series: <ChartSeries<PriceData, String>>[
                        LineSeries<PriceData, String>(
                          dataSource: widget.data,
                          xValueMapper: (PriceData sales, _) => sales.time,
                          yValueMapper: (PriceData sales, _) => sales.value,
                        )
                      ]),
                ),
                Row(
                  mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                  children: [
                    Flexible(
                        flex: 1,
                        child: Container(
                          height: 50,
                          child: Text(
                            "Available Amount: " + availableAmount.toString(),
                            style: TextStyle(fontSize: 24),
                          ),
                        )),
                    Flexible(
                        flex: 1,
                        child: Container(
                          height: 50,
                          child: Text(
                            widget.currentPrice + " €",
                            style: TextStyle(fontSize: 24),
                          ),
                        )),
                  ],
                ),
                Row(
                  mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                  children: [
                    Flexible(
                      flex: 1,
                      child: Container(
                        child: TextFormField(
                          keyboardType: TextInputType.number,
                          controller: _controller,
                          decoration:
                              const InputDecoration(border: OutlineInputBorder()),
                          onChanged: (value) {
                            buyAmount = int.parse(value);
                          },
                        ),
                      ),
                    ),
                    Flexible(
                      flex: 1,
                      child: Container(
                        height: 50,
                        width: 300,
                        child: RaisedButton(
                          color: Colors.red,
                          child: Text(
                            "Buy",
                            style: TextStyle(color: Colors.white),
                          ),
                          onPressed: () {
                            showDialog(
                                context: context,
                                builder: (context) {
                                  return AlertDialog(
                                    title: Text("Success"),
                                    content: Text("Order Completed."),
                                    actions: [
                                      TextButton(
                                          onPressed: () {
                                            availableAmount =
                                                availableAmount - buyAmount;
                                            setState(() {});
                                            Navigator.pop(context);
                                          },
                                          child: Text("Close"))
                                    ],
                                  );
                                });
                          },
                        ),
                      ),
                    ),
                  ],
                ),
                SizedBox(height: 10),
                Row(
                  mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                  children: [
                    Flexible(
                      flex: 1,
                      child: Container(
                        child: TextFormField(
                          keyboardType: TextInputType.number,
                          controller: _sellController,
                          decoration:
                              const InputDecoration(border: OutlineInputBorder()),
                          onChanged: (value) {
                            sellAmount = int.parse(value);
                          },
                        ),
                      ),
                    ),
                    Flexible(
                      flex: 1,
                      child: Container(
                        height: 50,
                        width: 300,
                        child: RaisedButton(
                          color: Colors.red,
                          child: Text(
                            "Sell",
                            style: TextStyle(color: Colors.white),
                          ),
                          onPressed: () {
                            showDialog(
                                context: context,
                                builder: (context) {
                                  return AlertDialog(
                                    title: Text("Success"),
                                    content: Text("Order Completed."),
                                    actions: [
                                      TextButton(
                                          onPressed: () {
                                            availableAmount =
                                                availableAmount + sellAmount;
                                            setState(() {});
                                            Navigator.pop(context);
                                          },
                                          child: Text("Close"))
                                    ],
                                  );
                                });
                          },
                        ),
                      ),
                    ),
                  ],
                )
              ],
            ),
          ),
        ),
      ),
    );
  }

  createDropDownItems() {
    List<DropdownMenuItem<String>> list = [];
    cities.forEach((element) {
      var item = DropdownMenuItem<String>(
          value: element,
          child: Container(
            child: Center(
                child: Text(
              element,
              style: TextStyle(fontSize: 30, color: Colors.black),
            )),
          ));
      list.add(item);
    });
    return list;
  }
}

class _SalesData {
  _SalesData(this.year, this.value);

  final int year;
  final int value;
}
