import type { ColorCollection } from '../colors';
import type { IResponse } from './IResponse';

export interface IUnspentOutput {
  outputID: {
    base58: string;
    transactionID: string;
    outputIndex: number;
  };

  type: string;

  output: {
    balances: {
      [color: string]: bigint;
    };

    address: string;
  };

  inclusionState: {
    confirmed?: boolean;
    rejected?: boolean;
    conflicting?: boolean;
  };
}

export interface IUnspentOutputAddress {
  type: string;
  base58: string;
}

export interface IUnspentOutputMap {
  address: string;
  outputs: {
    id: string;
    balances: ColorCollection;
    inclusionState: {
      confirmed?: boolean;
      rejected?: boolean;
      conflicting?: boolean;
    };
  }[];
}

export interface IUnspentOutputsResponse extends IResponse {
  unspentOutputs: {
    address: IUnspentOutputAddress;
    outputs: {
      output: IUnspentOutput;

      inclusionState: {
        confirmed?: boolean;
        rejected?: boolean;
        conflicting?: boolean;
      };
    }[];
  }[];
}

export interface ISingleUnspentOutputResponse extends IResponse {
  address: IUnspentOutputAddress;
  outputs: IUnspentOutput[];
}
