import { TestBed, inject } from '@angular/core/testing';

import { GoNewsApiService } from './go-news-api.service';

describe('GoNewsApiService', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [GoNewsApiService]
    });
  });

  it('should be created', inject([GoNewsApiService], (service: GoNewsApiService) => {
    expect(service).toBeTruthy();
  }));
});
