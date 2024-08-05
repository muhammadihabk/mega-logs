import { Injectable } from '@nestjs/common';
import * as fs from 'fs';
import * as Papa from 'papaparse';
import * as bluebird from 'bluebird';
import axios from 'axios';

@Injectable()
export class GenerateDataService {
  private startRow: number;
  private fixedProductFileName: string;

  constructor() {
    this.startRow = 0;
    this.fixedProductFileName = 'new_product';
  }

  async fixProductsFile(): Promise<void> {
    let newDataPapaParseFormat: string[] =
      await this.generateNewTranslatedFile();
    newDataPapaParseFormat.splice(0, 1)[0];
    let fields = [
      'product_key',
      'product_category_name',
      'product_name_length',
      'product_description_length',
      'product_photos_qty',
      'product_weight_g',
      'product_length_cm',
      'product_height_cm',
      'product_width_cm',
    ];
    const newFileData = Papa.unparse({
      fields,
      data: newDataPapaParseFormat,
    });

    fs.writeFileSync(`data/${this.fixedProductFileName}.csv`, newFileData);
  }

  getTranslationMap() {
    return new Promise((resolve, reject) => {
      let translation = {};
      const translationFilePath = 'data/product_category_name_translation.csv';
      const readStream = fs.createReadStream(translationFilePath, 'utf8');
      const config = {
        step: (results) => {
          translation[results.data[0]] = results.data[1];
        },
        complete: () => {
          resolve(translation);
        },
        error: (error) => {
          reject(error);
        },
      };
      Papa.parse(readStream, config);
    });
  }

  generateNewTranslatedFile(): Promise<string[]> {
    return new Promise(async (resolve, reject) => {
      const translationMap = await this.getTranslationMap();
      const productsFilePath = 'data/products.csv';
      const readStream = fs.createReadStream(productsFilePath, 'utf8');
      const config = {
        transform: (value, field) => {
          if (field == 1) {
            return translationMap[value];
          } else {
            return value;
          }
        },
        complete: (results) => {
          resolve(results.data);
        },
        error: (error) => {
          reject(error);
        },
      };
      Papa.parse(readStream, config);
    });
  }

  async parseRows(filePath, numRows): Promise<string[]> {
    return new Promise((resolve, reject) => {
      const rows = [];
      let rowCounter = 0;

      const stream = fs.createReadStream(filePath, 'utf8');

      const csvConfig = {
        step: (results, parser) => {
          rowCounter++;
          if (
            rowCounter > this.startRow &&
            rowCounter <= this.startRow + numRows
          ) {
            rows.push(results.data);
          }
          if (rows.length === numRows) {
            parser.abort();
            stream.close();
            resolve(rows);
          }
        },
        complete: () => {
          resolve(rows);
        },
        error: (error) => {
          reject(error);
        },
      };

      Papa.parse(stream, csvConfig);
    });
  }

  async getEntities() {
    const entities = {
      customers: { filePath: 'data/customers.csv', data: [] },
      products: { filePath: `data/${this.fixedProductFileName}.csv`, data: [] },
      orders: { filePath: 'data/orders.csv', data: [] },
      orderItems: { filePath: 'data/order_items.csv', data: [] },
      sellers: { filePath: 'data/sellers.csv', data: [] },
    };
    const numRows = 1000;

    let customers = await this.parseRows(entities.customers.filePath, numRows);
    entities.customers.data = customers.map((customer) => {
      return {
        customer_key: customer[0],
        customer_zip_code_prefix: customer[2],
        customer_city: customer[3],
        customer_state: customer[4],
      };
    });
    let orderItems = await this.parseRows(
      entities.orderItems.filePath,
      numRows,
    );
    let products = await this.parseRows(entities.products.filePath, numRows);
    entities.products.data = products.map((product) => {
      return {
        product_key: product[0],
        product_category_name: product[1],
        product_name_lenght: product[2],
        product_description_lenght: product[3],
        product_photos_qty: product[4],
        product_weight_g: product[5],
        product_length_cm: product[6],
        product_height_cm: product[7],
        product_width_cm: product[8],
      };
    });
    let orders = await this.parseRows(entities.orders.filePath, numRows);
    entities.orders.data = orders.map((order) => {
      return {
        order_key: order[0],
        customer_key: order[1],
        order_status: order[2],
        order_purchase_timestamp: order[3],
        order_approved_at: order[4],
        order_delivered_carrier_date: order[5],
        order_delivered_customer_date: order[6],
        order_estimated_delivery_date: order[7],
      };
    });
    entities.orderItems.data = orderItems.map((orderItem) => {
      return {
        order_key: orderItem[0],
        order_item_num: orderItem[1],
        product_key: orderItem[2],
        seller_id: orderItem[3],
        shipping_limit_date: orderItem[4],
        price: orderItem[5],
        freight_value: orderItem[6],
      };
    });
    let sellers = await this.parseRows(entities.sellers.filePath, numRows);
    entities.sellers.data = sellers.map((seller) => {
      return {
        seller_key: seller[0],
        seller_zip_code_prefix: seller[1],
        seller_city: seller[2],
        seller_state: seller[3],
      };
    });
    this.startRow += numRows;

    return entities;
  }

  async sendDataToGateway() {
    try {
      const entities = await this.getEntities();
      const promises = [
        entities.customers.data.map((element) => {
          axios.post('http://gateway:3000/route-requests/customers', element);
        }),
        entities.products.data.map((element) => {
          axios.post('http://gateway:3000/route-requests/products', element);
        }),
        entities.orders.data.map((element) => {
          axios.post('http://gateway:3000/route-requests/orders', element);
        }),
        entities.orderItems.data.map((element) => {
          axios.post('http://gateway:3000/route-requests/orderItems', element);
        }),
        entities.sellers.data.map((element) => {
          axios.post('http://gateway:3000/route-requests/sellers', element);
        }),
      ];
      bluebird.all(promises);
    } catch (error) {
      throw new Error(error);
    }
  }
}
