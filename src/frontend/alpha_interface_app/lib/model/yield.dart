class Yield {
  String domain = "";
  String area = "";
  String element = "";
  String item = "";
  int year = 0;
  String unit = "";
  int value = 0;

  Yield(this.domain, this.area, this.element, this.item, this.year, this.unit,
      this.value);

  Yield.fromJson(Map<String, dynamic> json) {
    domain = json['Domain'];
    area = json['Area'];
    element = json['Element'];
    item = json['Item'];
    year = json['Year'];
    unit = json['Unit'];
    value = json['Value'];
  }
}
